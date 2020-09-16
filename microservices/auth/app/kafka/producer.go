package kafka

import (
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func newProducer() (*kafka.Producer, error) {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": os.Getenv("KAFKA_URL")})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return producer, nil
}

func produceMessageReport(producer *kafka.Producer) {
	for e := range producer.Events() {
		switch ev := e.(type) {
		case *kafka.Message:
			if ev.TopicPartition.Error != nil {
				log.Printf("Delivery failed: %v\n", ev.TopicPartition)
			} else {
				log.Printf("Delivered message to %+v\n", ev.TopicPartition)
			}
		default:
			log.Printf("Kafka Error: %v\n", e)
		}
	}
}
