package data

import (
	"encoding/json"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// KafkaUpdatePassword is the JSON struct sent by the auth MS using the kafka topic `user.password.update`.
type KafkaUpdatePassword struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}

// GetUserUpdatePasswordMessage return a KafkaUpdatePassword from a kafka message.
func GetUserUpdatePasswordMessage(msg kafka.Message) (KafkaUpdatePassword, error) {
	var userLoginMessage KafkaUpdatePassword
	if err := json.Unmarshal(msg.Value, &userLoginMessage); err != nil {
		log.Printf("Topic: %s -> error getting message from string: %s\nerror: %s",
			msg.TopicPartition, string(msg.Value), err)
		return KafkaUpdatePassword{}, err
	}
	return userLoginMessage, nil
}
