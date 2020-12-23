package producers

import (
	"log"
	"time"

	"github.com/alexandr-io/backend/user/data"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// produceUserResponse produce a message to the `user-response` topic.
func produceUserResponse(key string, message []byte) error {
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
		TopicPartition: kafka.TopicPartition{Topic: &userResponse.Name, Partition: kafka.PartitionAny},
		Key:            []byte(key),
		Value:          message,
	}, nil); err != nil {
		log.Println(err)
		return err
	}

	// Wait for message deliveries before shutting down
	producer.Flush(int((15 * time.Microsecond).Microseconds()))
	return nil
}

// SendSuccessUserMessage create a success user response and send it the the topic.
func SendSuccessUserMessage(key string, code int, user data.User) error {
	message, err := data.CreateUserResponseMessage(code,
		data.KafkaUser{
			ID:       user.ID,
			Email:    user.Email,
			Username: user.Username,
		})
	if err != nil {
		return err
	}

	return produceUserResponse(key, message)
}

// SendInternalErrorUserMessage create an internal error response and send it the the `user-response` topic.
func SendInternalErrorUserMessage(key string, content string) error {
	message, err := data.CreateKafkaInternalErrorMessage(content)
	if err != nil {
		return err
	}

	return produceUserResponse(key, message)
}

// SendUnauthorizedUserMessage create an bad request error response and send it the the `user-response` topic.
func SendUnauthorizedUserMessage(key string, content string) error {
	message, err := data.CreateKafkaUnauthorizedErrorMessage(content)
	if err != nil {
		return err
	}

	return produceUserResponse(key, message)
}
