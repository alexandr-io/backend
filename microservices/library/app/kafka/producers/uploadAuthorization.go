package producers

import (
	"log"
	"time"

	"github.com/alexandr-io/backend/library/data"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// produceLibraryUploadAuthorizationResponse produce a message to the `library.upload.allowed.response` topic.
func produceLibraryUploadAuthorizationResponse(key string, message []byte) error {
	// Create a new producer
	producer, err := newProducer()
	if err != nil {
		return err
	}
	defer producer.Close()

	// Delivery report handler for produced messages
	go produceMessageReport(producer)

	// Produce message to topic (asynchronously)
	if err := producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &libraryUploadAllowed.Name, Partition: kafka.PartitionAny},
		Key:            []byte(key),
		Value:          message,
	}, nil); err != nil {
		log.Println(err)
		return err
	}

	// Wait for message deliveries before shutting down
	producer.Flush(int((5 * time.Microsecond).Microseconds()))
	return nil
}

// SendInternalErrorLibraryUploadAuthorizationMessage create an internal error response and send it the the `library.upload.allowed.response` topic.
func SendInternalErrorLibraryUploadAuthorizationMessage(key string, content string) error {
	message, err := data.CreateKafkaInternalErrorMessage(content)
	if err != nil {
		return err
	}

	return produceLibraryUploadAuthorizationResponse(key, message)
}

// SendSuccessLibraryUploadAuthorizationMessage create a success library authorization response and send it the the topic.
func SendSuccessLibraryUploadAuthorizationMessage(key string, code int, isAllowed bool) error {
	message, err := data.CreateLibraryUploadAuthorizationResponseMessage(code, data.KafkaLibraryUploadAuthorizationContent{
		IsAllowed: isAllowed,
	})
	if err != nil {
		return err
	}

	return produceLibraryUploadAuthorizationResponse(key, message)
}
