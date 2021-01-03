package producers

import (
	"log"
	"time"

	"github.com/alexandr-io/backend/user/data"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// produceUpdatePasswordResponse produce a message to the `user.password.update.response` topic.
func produceUpdatePasswordResponse(key string, message []byte) error {
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
		TopicPartition: kafka.TopicPartition{Topic: &updatePasswordResponse.Name, Partition: kafka.PartitionAny},
		Key:            []byte(key),
		Value:          message,
	}, nil); err != nil {
		log.Println(err)
		return err
	}

	// Wait for message deliveries before shutting down
	producer.Flush(int((time.Second).Microseconds()))
	return nil
}

// SendSuccessUpdatePasswordMessage create a success updatePassword response and send it the the topic.
func SendSuccessUpdatePasswordMessage(key string, code int, user data.User) error {
	message, err := data.CreateUserResponseMessage(code,
		data.KafkaUser{
			ID:       user.ID,
			Email:    user.Email,
			Username: user.Username,
		})
	if err != nil {
		return err
	}

	return produceUpdatePasswordResponse(key, message)
}

// SendInternalErrorUpdatePasswordMessage create an internal error response and send it the the `user.password.update.response` topic.
func SendInternalErrorUpdatePasswordMessage(key string, content string) error {
	message, err := data.CreateKafkaInternalErrorMessage(content)
	if err != nil {
		return err
	}

	return produceUpdatePasswordResponse(key, message)
}
