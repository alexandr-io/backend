package producers

import (
	"encoding/json"
	"github.com/alexandr-io/berrors"
	"log"
	"time"

	"github.com/alexandr-io/backend/user/data"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func ProduceRegisterResponseMessage(id string, messageData []byte) error {
	// Create the message in the correct format
	message, err := data.CreateMessage(id, messageData)
	if err != nil {
		return err
	}

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

func SendKafkaMessageToProducer(ID string, kafkaMessage berrors.KafkaErrorMessage) {
	kafkaMessageJSON, err := json.Marshal(kafkaMessage)
	if err != nil {
		log.Println(err)
		return
	}
	if err := ProduceRegisterResponseMessage(ID, kafkaMessageJSON); err != nil {
		log.Println(err)
		return
	}
}
