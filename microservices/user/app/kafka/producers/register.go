package producers

import (
	"encoding/json"
	"log"
	"time"

	"github.com/alexandr-io/backend/user/data"
	"github.com/alexandr-io/berrors"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// produceRegisterResponse produce a message to the `register-response` topic.
func produceRegisterResponse(message []byte) error {
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
		TopicPartition: kafka.TopicPartition{Topic: &registerResponse, Partition: kafka.PartitionAny},
		Value:          message,
	}, nil); err != nil {
		log.Println(err)
		return err
	}

	// Wait for message deliveries before shutting down
	producer.Flush(int((15 * time.Microsecond).Microseconds()))
	return nil
}

// SendSuccessRegisterMessage create a success register response and send it the the topic.
func SendSuccessRegisterMessage(id string, code int, user data.User) error {
	message, err := data.CreateRegisterResponseMessage(id, code,
		data.KafkaUserRegisterResponseContent{
			Email:    user.Email,
			Username: user.Username,
		})
	if err != nil {
		return err
	}

	return produceRegisterResponse(message)
}

// SendInternalErrorRegisterMessage create an internal error response and send it the the `register-response` topic.
func SendInternalErrorRegisterMessage(id string, content string) error {
	message, err := data.CreateKafkaInternalErrorMessage(id, content)
	if err != nil {
		return err
	}

	return produceRegisterResponse(message)
}

// SendBadRequestRegisterMessage create an bad request error response and send it the the `register-response` topic.
func SendBadRequestRegisterMessage(id string, content []byte) error {
	var badInput berrors.BadInput
	if err := json.Unmarshal(content, &badInput); err != nil {
		return err
	}

	message, err := data.CreateKafkaBadRequestMessage(id, badInput)
	if err != nil {
		return err
	}

	return produceRegisterResponse(message)
}