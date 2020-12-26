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

// UserRequestChannels is the map of channel used to store uuid of kafka message.
// This is made to retrieve the response message corresponding to the request.
var UserRequestChannels sync.Map

// UserRequestHandler is the entry point of a new user message to the user MS using kafka.
// The function produce a new message to kafka,
// create a channel for the answer,
// call a watcher to wait for the proper answer from the user-response topic,
// interpret the answer (possible errors or success) and return and error with the proper http code
// In case of success, a data.User is returned containing the username and email of the user.
func UserRequestHandler(user data.KafkaUser) (*data.KafkaUser, error) {
	// Generate UUID
	id := uuid.New()
	// Create a channel for the request
	requestChannel := make(chan string)
	// Create a channel to manage error in goroutines
	errorChannel := make(chan error)
	// Store request channel to the channel map
	UserRequestChannels.Store(id.String(), requestChannel)
	// Produce the message to kafka
	message, err := user.MarshalKafkaUser()
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	go produceUserMessage(id.String(), string(message), errorChannel)
	// Watch for a response in the request channel
	kafkaMessage, rawMessage, err := userResponseWatcher(id.String(), requestChannel, errorChannel)
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

// produceUserMessage produce a user message to the `user` topic.
// The message sent is the userID.
func produceUserMessage(id string, user string, errorChannel chan error) {
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
		TopicPartition: kafka.TopicPartition{Topic: &userRequest.Name, Partition: kafka.PartitionAny},
		Key:            []byte(id),
		Value:          []byte(user),
	}, nil); err != nil {
		errorChannel <- data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
		return
	}

	// Wait for message deliveries before shutting down
	producer.Flush(int((15 * time.Microsecond).Microseconds()))
	return
}

// userResponseWatcher is waiting for an update in the given channel. The message will be set in the channel by
// consumeUserResponseMessages once the user MS has answered to the request.
// The channel is then deleted from the map and the kafka message is returned.
func userResponseWatcher(id string, requestChannel chan string, errorChannel chan error) (*data.KafkaResponseMessage, []byte, error) {
	timeout := time.After(5 * time.Second)
	for {
		select {
		case <-timeout:
			// In case of time out, we delete the channel and return an error
			UserRequestChannels.Delete(id)
			return nil, nil, data.NewHTTPErrorInfo(fiber.StatusGatewayTimeout, "Kafka user response timed out")
		case err := <-errorChannel:
			return nil, nil, err
		case message := <-requestChannel:
			var kafkaMessage data.KafkaResponseMessage
			if err := json.Unmarshal([]byte(message), &kafkaMessage); err != nil {
				return nil, nil, err
			}
			UserRequestChannels.Delete(id)
			return &kafkaMessage, []byte(message), nil
		}
	}
}
