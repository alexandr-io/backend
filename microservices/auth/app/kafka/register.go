package kafka

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/alexandr-io/backend/auth/data"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/google/uuid"
)

var registerRequestChannels sync.Map

func RegisterRequestHandler(user data.UserRegister) error {
	// Generate UUID
	id := uuid.New()
	// Create a channel for the request
	requestChannel := make(chan string)
	// Channel to manage error in goroutines
	errorChannel := make(chan error)
	// Store request channel to the channel map
	registerRequestChannels.Store(id.String(), requestChannel)
	// Produce the message to kafka
	go produceRegisterMessage(id.String(), user, errorChannel)
	// Watch for a response in the request channel
	if err := registerResponseWatcher(id.String(), requestChannel, errorChannel); err != nil {
		return err
	}
	return nil
}

func produceRegisterMessage(id string, user data.UserRegister, errorChannel chan error) {
	// Create the message in the correct format
	user.ConfirmPassword = "" // Not needed
	message, err := createMessage(id, user)
	if err != nil {
		errorChannel <- err
		return
	}

	// Create a new producer
	producer, err := newProducer()
	if err != nil {
		log.Println(err)
		errorChannel <- err
		return
	}
	defer producer.Close()

	// Delivery report handler for produced messages
	go produceMessageReport(producer)

	// Produce message to topic (asynchronously)
	if err := producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &register, Partition: kafka.PartitionAny},
		Value:          message,
	}, nil); err != nil {
		log.Println(err)
		errorChannel <- err
		return
	}

	// Wait for message deliveries before shutting down
	producer.Flush(int((15 * time.Microsecond).Microseconds()))
	return
}

func consumeRegisterResponseMessages() {
	// Create new consumer
	consumer, err := newConsumer()
	if err != nil {
		return
	}
	defer consumer.Close()

	// Subscribe consumer to topic register-
	if err := consumer.SubscribeTopics([]string{registerResponse}, nil); err != nil {
		log.Println(err)
		return
	}

	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
			messageData, err := getMessageFromBytes(msg.Value)
			if err != nil {
				continue
			}
			requestChannelInterface, ok := registerRequestChannels.Load(messageData.UUID)
			if !ok {
				log.Printf("Can't load channel %s from requestChannels", messageData.UUID)
				continue
			}
			requestChannel := requestChannelInterface.(chan string)
			requestChannel <- string(messageData.Data)
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}

func registerResponseWatcher(id string, requestChannel chan string, errorChannel chan error) error {
	for {
		fmt.Println("IN")
		select {
		case err := <-errorChannel:
			return err
		case message := <-requestChannel:
			fmt.Println(message)
			registerRequestChannels.Delete(id)
			return nil
		}
	}
}
