package producers

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/alexandr-io/backend/library/data"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// AuthRequestChannels is the map of channel used to store uuid of kafka message.
// This is made to retrieve the response message corresponding to the request.
var AuthRequestChannels sync.Map

// AuthRequestHandler is the entry point of a new auth message to the auth MS using kafka.
// The function produce a new message to kafka,
// create a channel for the answer,
// call a watcher to wait for the proper answer from the auth-response topic,
// interpret the answer (possible errors or success) and return and error with the proper http code
// In case of success, a data.KafkaUser is returned containing the id, username and email of the user.
func AuthRequestHandler(jwt string) (*data.KafkaUserResponseContent, error) {
	// Generate UUID
	id := uuid.New()
	// Create a channel for the request
	requestChannel := make(chan string)
	// Create a channel to manage error in goroutines
	errorChannel := make(chan error)
	// Store request channel to the channel map
	AuthRequestChannels.Store(id.String(), requestChannel)
	// Produce the message to kafka
	go produceAuthMessage(id.String(), jwt, errorChannel)
	// Watch for a response in the request channel
	kafkaMessage, rawMessage, err := authResponseWatcher(id.String(), requestChannel, errorChannel)
	if err != nil {
		return nil, err
	}

	// handle error
	if err := handleError(*kafkaMessage, rawMessage); err != nil {
		return nil, err
	}

	// handle success
	if kafkaMessage.Code == fiber.StatusOK {
		return data.UnmarshalAuthResponse(rawMessage)
	}

	// If http code contained in the kafka message is not handled
	return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, fmt.Sprintf("unmanaged code: %d", kafkaMessage.Code))
}

// produceAuthMessage produce a auth message to the `auth` topic.
// The message sent is a JSON of the jwt.
func produceAuthMessage(id string, jwt string, errorChannel chan error) {
	message, err := data.CreateAuthRequestMessage(jwt)
	if err != nil {
		errorChannel <- data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
		return
	}

	// Create a new producers
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
		TopicPartition: kafka.TopicPartition{Topic: &authRequest.Name, Partition: kafka.PartitionAny},
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

// authResponseWatcher is waiting for an update in the given channel. The message will be set in the channel by
// consumeAuthResponseMessages once the auth MS has answered to the request.
// The channel is then deleted from the map and the kafka message is returned.
func authResponseWatcher(id string, requestChannel chan string, errorChannel chan error) (*data.KafkaResponseMessage, []byte, error) {
	timeout := time.After(15 * time.Second)
	for {
		select {
		case <-timeout:
			// In case of time out, we delete the channel and return an error
			AuthRequestChannels.Delete(id)
			return nil, nil, data.NewHTTPErrorInfo(fiber.StatusGatewayTimeout, "Kafka auth response timed out")
		case err := <-errorChannel:
			return nil, nil, err
		case message := <-requestChannel:
			var kafkaMessage data.KafkaResponseMessage
			if err := json.Unmarshal([]byte(message), &kafkaMessage); err != nil {
				return nil, nil, err
			}
			AuthRequestChannels.Delete(id)
			return &kafkaMessage, []byte(message), nil
		}
	}
}
