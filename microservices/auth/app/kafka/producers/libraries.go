package producers

import (
	"time"

	"github.com/alexandr-io/backend/auth/data"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/google/uuid"
)

// CreateUserLibrariesRequestHandler produce a message to the `libraries-creation-request` topic.
func CreateUserLibrariesRequestHandler(user data.KafkaLibrariesCreationMessage) error {
	// Generate UUID
	id := uuid.New()

	return produceCreateLibraryMessage(id.String(), user)

}

func produceCreateLibraryMessage(id string, user data.KafkaLibrariesCreationMessage) error {
	message, err := data.CreateLibrariesCreationMessage(user)
	if err != nil {
		return err
	}

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
		TopicPartition: kafka.TopicPartition{Topic: &librariesRequest.Name, Partition: kafka.PartitionAny},
		Key:            []byte(id),
		Value:          message,
	}, nil); err != nil {
		return err
	}

	// Wait for message deliveries before shutting down
	producer.Flush(int((time.Microsecond * 50).Microseconds()))
	return nil
}
