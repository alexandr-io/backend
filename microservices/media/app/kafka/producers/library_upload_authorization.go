package producers

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/alexandr-io/backend/media/data"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// LibraryUploadAuthorizationRequestChannels is the map of channel used to store uuid of kafka message.
// This is made to retrieve the response message corresponding to the request.
var LibraryUploadAuthorizationRequestChannels sync.Map

// LibraryUploadAuthorizationRequestHandler is the entry point of a new library.upload.allowed message to the library MS using kafka.
// The function produce a new message to kafka,
// create a channel for the answer,
// call a watcher to wait for the proper answer from the library.upload.allowed.response topic,
// interpret the answer (possible errors or success) and return and error with the proper http code
// In case of success, a bool is returned telling if the user can upload a book to the library.
func LibraryUploadAuthorizationRequestHandler(book *data.Book, userID string) (bool, error) {
	// Generate UUID
	id := uuid.New()
	// Create a channel for the request
	requestChannel := make(chan string)
	// Create a channel to manage error in goroutines
	errorChannel := make(chan error)
	// Store request channel to the channel map
	LibraryUploadAuthorizationRequestChannels.Store(id.String(), requestChannel)
	// Produce the message to kafka
	go produceLibraryUploadAuthorizationMessage(id.String(), book, userID, errorChannel)
	// Watch for a response in the request channel
	kafkaMessage, rawMessage, err := libraryUploadAuthorizationResponseWatcher(id.String(), requestChannel, errorChannel)
	if err != nil {
		return false, err
	}

	// handle error
	if err := handleError(*kafkaMessage, rawMessage); err != nil {
		return false, err
	}

	// handle success
	if kafkaMessage.Code == fiber.StatusOK {
		response, err := data.UnmarshalLibraryAuthorizationResponse(rawMessage)
		if err != nil {
			return false, err
		}
		return response.IsAllowed, nil
	}

	// If http code contained in the kafka message is not handled
	return false, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, fmt.Sprintf("unmanaged code: %d", kafkaMessage.Code))
}

// produceLibraryUploadAuthorizationMessage produce a library authorization message to the `library.upload.allowed` topic.
// The message sent is a JSON of the book and user information.
func produceLibraryUploadAuthorizationMessage(id string, book *data.Book, userID string, errorChannel chan error) {
	message, err := data.CreateLibraryAuthorizationRequestMessage(book, userID)
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
		TopicPartition: kafka.TopicPartition{Topic: &libraryCanUpload.Name, Partition: kafka.PartitionAny},
		Key:            []byte(id),
		Value:          message,
	}, nil); err != nil {
		errorChannel <- data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
		return
	}

	// Wait for message deliveries before shutting down
	producer.Flush(int((time.Microsecond * 50).Microseconds()))
	return
}

// libraryUploadAuthorizationResponseWatcher is waiting for an update in the given channel. The message will be set in the channel by
// consumeAuthResponseMessages once the library MS has answered to the request.
// The channel is then deleted from the map and the kafka message is returned.
func libraryUploadAuthorizationResponseWatcher(id string, requestChannel chan string, errorChannel chan error) (*data.KafkaResponseMessage, []byte, error) {
	timeout := time.After(15 * time.Second)
	for {
		select {
		case <-timeout:
			// In case of time out, we delete the channel and return an error
			LibraryUploadAuthorizationRequestChannels.Delete(id)
			return nil, nil, data.NewHTTPErrorInfo(fiber.StatusGatewayTimeout, "Kafka library response timed out")
		case err := <-errorChannel:
			return nil, nil, err
		case message := <-requestChannel:
			var kafkaMessage data.KafkaResponseMessage
			if err := json.Unmarshal([]byte(message), &kafkaMessage); err != nil {
				return nil, nil, err
			}
			LibraryUploadAuthorizationRequestChannels.Delete(id)
			return &kafkaMessage, []byte(message), nil
		}
	}
}
