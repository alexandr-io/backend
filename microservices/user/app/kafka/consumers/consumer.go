package consumers

import (
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// StartConsumers starts all the kafka consumers in goroutines.
func StartConsumers() {
	go consumeRegisterRequestMessages()
	go consumeLoginRequestMessages()
	go consumeUserRequestMessages()
}

// newConsumer create a new kafka consumer.
func newConsumer() (*kafka.Consumer, error) {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_URL"),
		"group.id":          "group",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		log.Println(err)
		return nil, err
	}
	return consumer, nil
}
