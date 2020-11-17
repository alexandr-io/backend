package consumers

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"os"
)

// StartConsumers starts all the kafka consumers in goroutines.
func StartConsumers() {
	go consumeAuthResponseMessages()

	go consumeLibrariesCreationMessages()
}

// newConsumer create a new kafka consumer.
func newConsumer(consumerGroup string) (*kafka.Consumer, error) {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_URL"),
		"group.id":          consumerGroup,
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		log.Println(err)
		return nil, err
	}
	return consumer, nil
}
