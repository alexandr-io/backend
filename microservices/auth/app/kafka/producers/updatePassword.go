package producers

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/alexandr-io/backend/auth/data"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// UpdatePasswordRequestChannels is the map of channel used to store uuid of kafka message.
// This is made to retrieve the response message corresponding to the request.
var UpdatePasswordRequestChannels sync.Map

// UpdatePasswordRequestHandler is the entry point of a new user.update.password message to the user MS using kafka.
// The function produce a new message to kafka,
// create a channel for the answer,
// call a watcher to wait for the proper answer from the user.update.password.response topic,
// interpret the answer (possible errors or success) and return and error with the proper http code
// In case of success, a data.KafkaUser is returned containing the username and email of the user.
func UpdatePasswordRequestHandler(userUpdate data.KafkaUpdatePassword) (*data.KafkaUser, error) {
	// Generate UUID
	id := uuid.New()
	// Create a channel for the request
	requestChannel := make(chan string)
	// Create a channel to manage error in goroutines
	errorChannel := make(chan error)
	// Store request channel to the channel map
	UpdatePasswordRequestChannels.Store(id.String(), requestChannel)
	// Produce the message to kafka
	go produceUpdatePasswordMessage(id.String(), userUpdate, errorChannel)
	// Watch for a response in the request channel
	kafkaMessage, rawMessage, err := updatePasswordResponseWatcher(id.String(), requestChannel, errorChannel)
	if err != nil {
		return nil, err
	}

	// handle error
	if err := handleError(*kafkaMessage, rawMessage); err != nil {
		return nil, err
	}

	// handle success
	if kafkaMessage.Code == fiber.StatusOK {
		return data.UnmarshalUserResponse(rawMessage)
	}

	// If http code contained in the kafka message is not handled
	return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, fmt.Sprintf("unmanaged code: %d", kafkaMessage.Code))
}

// produceUpdatePasswordMessage produce a updatePassword message to the `updatePassword` topic.
// The message sent is a JSON of the data.KafkaUserUpdatePasswordMessage struct.
func produceUpdatePasswordMessage(id string, userUpdate data.KafkaUpdatePassword, errorChannel chan error) {
	message, err := userUpdate.Marshall()
	if err != nil {
		errorChannel <- data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
		return
	}

	// Create a new producer
	producer, err := newProducer()
	if err != nil {
		errorChannel <- data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
		return
	}
	defer producer.Close()

	// Delivery report handler for produced messages
	go produceMessageReport(producer)

	// Produce message to topic (asynchronously)
	if err := producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &updatePasswordRequest.Name, Partition: kafka.PartitionAny},
		Key:            []byte(id),
		Value:          message,
	}, nil); err != nil {
		errorChannel <- data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
		return
	}

	// Wait for message deliveries before shutting down
	producer.Flush(int((time.Second).Microseconds()))
	return
}

// updatePasswordResponseWatcher is waiting for an update in the given channel. The message will be set in the channel by
// consumeUpdatePasswordResponseMessages once the user MS has answered to the request.
// The channel is then deleted from the map and the kafka message is returned.
func updatePasswordResponseWatcher(id string, requestChannel chan string, errorChannel chan error) (*data.KafkaResponseMessage, []byte, error) {
	timeout := time.After(15 * time.Second)
	for {
		select {
		case <-timeout:
			// In case of time out, we delete the channel and return an error
			UpdatePasswordRequestChannels.Delete(id)
			return nil, nil, data.NewHTTPErrorInfo(fiber.StatusGatewayTimeout, "Kafka updatePassword response timed out")
		case err := <-errorChannel:
			return nil, nil, err
		case message := <-requestChannel:
			var kafkaMessage data.KafkaResponseMessage
			if err := json.Unmarshal([]byte(message), &kafkaMessage); err != nil {
				return nil, nil, err
			}
			UpdatePasswordRequestChannels.Delete(id)
			return &kafkaMessage, []byte(message), nil
		}
	}
}
