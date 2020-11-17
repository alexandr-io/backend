package producers

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"os"
)

// newProducer create a new kafka producers.
func newProducer() (*kafka.Producer, error) {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": os.Getenv("KAFKA_URL")})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return producer, nil
}

// produceMessageReport watch the production of a message to kafka and log possible errors.
// This function must be launched in a goroutine after the calling the produce function.
func produceMessageReport(producer *kafka.Producer) {
	for e := range producer.Events() {
		switch ev := e.(type) {
		case *kafka.Message:
			if ev.TopicPartition.Error != nil {
				log.Printf("Delivery failed: %v\n", ev.TopicPartition)
			} else {
				fmt.Printf("[KAFKA]: Delivered message to %v\n", ev.TopicPartition)
			}
		default:
			log.Printf("Kafka Error: %v\n", e)
		}
	}
}
