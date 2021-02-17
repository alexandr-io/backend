package data

import (
	"encoding/json"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// KafkaAuthRequest is the data sent to the auth MS to check the validity of a jwt.
type KafkaAuthRequest struct {
	JWT string `json:"jwt"`
}

// GetAuthMessage return a KafkaAuthRequest from a kafka message.
func GetAuthMessage(msg kafka.Message) (KafkaAuthRequest, error) {
	var authMessage KafkaAuthRequest
	if err := json.Unmarshal(msg.Value, &authMessage); err != nil {
		log.Printf("Topic: %s -> error getting message from string: %s\nerror: %s",
			msg.TopicPartition, string(msg.Value), err)
		return KafkaAuthRequest{}, err
	}
	return authMessage, nil
}
