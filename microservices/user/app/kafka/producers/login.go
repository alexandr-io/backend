package producers

import (
	"encoding/json"
	"log"
	"time"

	"github.com/alexandr-io/backend/user/data"
	"github.com/alexandr-io/berrors"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// produceLoginResponse produce a message to the `login-response` topic.
func produceLoginResponse(key string, message []byte) error {
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
		TopicPartition: kafka.TopicPartition{Topic: &loginResponse.Name, Partition: kafka.PartitionAny},
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

// SendSuccessLoginMessage create a success login response and send it the the topic.
func SendSuccessLoginMessage(key string, code int, user data.User) error {
	message, err := data.CreateUserResponseMessage(code,
		data.KafkaUserResponseContent{
			ID:       user.ID,
			Email:    user.Email,
			Username: user.Username,
		})
	if err != nil {
		return err
	}

	return produceLoginResponse(key, message)
}

// SendInternalErrorLoginMessage create an internal error response and send it the the `login-response` topic.
func SendInternalErrorLoginMessage(key string, content string) error {
	message, err := data.CreateKafkaInternalErrorMessage(content)
	if err != nil {
		return err
	}

	return produceLoginResponse(key, message)
}

// SendBadRequestLoginMessage create an bad request error response and send it the the `login-response` topic.
func SendBadRequestLoginMessage(key string, content []byte) error {
	var badInput berrors.BadInput
	if err := json.Unmarshal(content, &badInput); err != nil {
		return err
	}

	message, err := data.CreateKafkaBadRequestMessage(badInput)
	if err != nil {
		return err
	}

	return produceLoginResponse(key, message)
}
