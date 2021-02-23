package consumers

import (
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// StartConsumers starts all the kafka consumers in goroutines.
func StartConsumers() {
	// consumers for backward communication
	go consumeRegisterResponseMessages()
	go consumeLoginResponseMessages()
	go consumeUpdatePasswordResponseMessages()
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
