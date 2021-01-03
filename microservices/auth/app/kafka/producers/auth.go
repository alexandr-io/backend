package producers

import (
	"log"
	"time"

	"github.com/alexandr-io/backend/auth/data"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// produceAuthResponse produce a message to the `auth-response` topic.
func produceAuthResponse(key string, message []byte) error {
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
		TopicPartition: kafka.TopicPartition{Topic: &authResponse.Name, Partition: kafka.PartitionAny},
		Key:            []byte(key),
		Value:          message,
	}, nil); err != nil {
		log.Println(err)
		return err
	}

	// Wait for message deliveries before shutting down
	producer.Flush(int((time.Microsecond * 50).Microseconds()))
	return nil
}

// SendAuthResponseMessage create a auth response and send it the the topic.
func SendAuthResponseMessage(key string, code int, user *data.User) error {
	message, err := data.CreateAuthResponseMessage(code,
		data.KafkaUser{
			ID:       user.ID,
			Email:    user.Email,
			Username: user.Username,
		})
	if err != nil {
		return err
	}

	return produceAuthResponse(key, message)
}

// SendErrorAuthMessage create an error response and send it the the `auth-response` topic.
func SendErrorAuthMessage(key string, err error) error {
	message, err := data.CreateKafkaErrorMessage(err)
	if err != nil {
		return err
	}

	return produceAuthResponse(key, message)
}
