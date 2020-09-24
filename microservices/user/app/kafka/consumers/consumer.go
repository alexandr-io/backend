package consumers

import (
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func StartConsumers() {
	go consumeRegisterRequestMessages()
}

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
