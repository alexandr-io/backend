package producers

import (
	"time"

	"github.com/alexandr-io/backend/auth/data"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// EmailRequestHandler produce a message to the `email.new` topic.
func EmailRequestHandler(emailData data.KafkaEmail) error {
	// Generate UUID
	id := uuid.New()
	// Produce the message to kafka
	message, err := emailData.Marshal()
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	producer, err := newProducer()
	if err != nil {
		return err
	}
	defer producer.Close()

	// Delivery report handler for produced messages
	go produceMessageReport(producer)

	// Produce message to topic (asynchronously)
	if err := producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &emailNew.Name, Partition: kafka.PartitionAny},
		Key:            []byte(id.String()),
		Value:          message,
	}, nil); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	// Wait for message deliveries before shutting down
	producer.Flush(int((time.Second).Microseconds()))
	return nil
}
