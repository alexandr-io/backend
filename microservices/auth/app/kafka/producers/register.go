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

// RegisterRequestChannels is the map of channel used to store uuid of kafka message.
// This is made to retrieve the response message corresponding to the request.
var RegisterRequestChannels sync.Map

// RegisterRequestHandler is the entry point of a new register message to the user MS using kafka.
// The function produce a new message to kafka,
// create a channel for the answer,
// call a watcher to wait for the proper answer from the register-response topic,
// interpret the answer (possible errors or success) and return and error with the proper http code
// In case of success, a data.User is returned containing the username and email of the new user.
func RegisterRequestHandler(user data.UserRegister) (*data.KafkaUser, error) {
	// Generate UUID
	id := uuid.New()
	// Create a channel for the request
	requestChannel := make(chan string)
	// Create a channel to manage error in goroutines
	errorChannel := make(chan error)
	// Store request channel to the channel map
	RegisterRequestChannels.Store(id.String(), requestChannel)
	// Produce the message to kafka
	go produceRegisterMessage(id.String(), user, errorChannel)
	// Watch for a response in the request channel
	kafkaMessage, rawMessage, err := registerResponseWatcher(id.String(), requestChannel, errorChannel)
	if err != nil {
		return nil, err
	}

	// handle error
	if err := handleError(*kafkaMessage, rawMessage); err != nil {
		return nil, err
	}

	// handle success
	if kafkaMessage.Code == fiber.StatusCreated {
		return data.UnmarshalUserResponse(rawMessage)
	}

	// If http code contained in the kafka message is not handled
	return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, fmt.Sprintf("unmanaged code: %d", kafkaMessage.Code))
}

// produceRegisterMessage produce a register message to the `register` topic.
// The message sent is a JSON of the data.KafkaUserRegisterMessage struct.
func produceRegisterMessage(id string, user data.UserRegister, errorChannel chan error) {
	// Create the message in the correct format
	user.ConfirmPassword = "" // Not needed
	message, err := data.CreateRegisterMessage(user)
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
		TopicPartition: kafka.TopicPartition{Topic: &registerRequest, Partition: kafka.PartitionAny},
		Key:            []byte(id),
		Value:          message,
	}, nil); err != nil {
		errorChannel <- data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
		return
	}

	// Wait for message deliveries before shutting down
	producer.Flush(int((15 * time.Microsecond).Microseconds()))
	return
}

// registerResponseWatcher is waiting for an update in the given channel. The message will be set in the channel by
// consumeRegisterResponseMessages once the user MS has answered to the request.
// The channel is then deleted from the map and the kafka message is returned.
func registerResponseWatcher(id string, requestChannel chan string, errorChannel chan error) (*data.KafkaResponseMessage, []byte, error) {
	timeout := time.After(5 * time.Second)
	for {
		select {
		case <-timeout:
			// In case of time out, we delete the channel and return an error
			RegisterRequestChannels.Delete(id)
			return nil, nil, data.NewHTTPErrorInfo(fiber.StatusGatewayTimeout, "Kafka register response timed out")
		case err := <-errorChannel:
			return nil, nil, err
		case message := <-requestChannel:
			var kafkaMessage data.KafkaResponseMessage
			if err := json.Unmarshal([]byte(message), &kafkaMessage); err != nil {
				return nil, nil, err
			}
			RegisterRequestChannels.Delete(id)
			return &kafkaMessage, []byte(message), nil
		}
	}
}
