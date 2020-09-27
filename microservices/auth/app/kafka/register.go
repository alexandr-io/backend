package kafka

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/alexandr-io/backend/auth/data"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gofiber/fiber"
	"github.com/google/uuid"
)

var registerRequestChannels sync.Map

func RegisterRequestHandler(ctx *fiber.Ctx, user data.UserRegister) (*data.User, error) {
	// Generate UUID
	id := uuid.New()
	// Create a channel for the request
	requestChannel := make(chan string)
	// Create a channel to manage error in goroutines
	errorChannel := make(chan error)
	// Store request channel to the channel map
	registerRequestChannels.Store(id.String(), requestChannel)
	// Produce the message to kafka
	go produceRegisterMessage(id.String(), user, errorChannel)
	// Watch for a response in the request channel
	kafkaMessage, rawMessage, err := registerResponseWatcher(id.String(), requestChannel, errorChannel)
	if err != nil {
		ctx.SendStatus(http.StatusInternalServerError)
		return nil, err
	}

	// handle error
	if errorSet := handleError(ctx, *kafkaMessage, rawMessage); errorSet {
		// So that the proper ctx error is set in register route
		return nil, errors.New("")
	}

	// handle success
	if kafkaMessage.Data.Code == http.StatusCreated {
		return data.UnmarshalRegisterResponse(rawMessage)
	}

	// If http code contained in the kafka message is not handled
	log.Printf("Unmanaged code: %d\n", kafkaMessage.Data.Code)
	ctx.SendStatus(http.StatusInternalServerError)
	return nil, fmt.Errorf("unmanaged code: %d\n", kafkaMessage.Data.Code)
}

func produceRegisterMessage(id string, user data.UserRegister, errorChannel chan error) {
	// Create the message in the correct format
	user.ConfirmPassword = "" // Not needed
	message, err := data.CreateRegisterMessage(id, user)
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
		TopicPartition: kafka.TopicPartition{Topic: &registerRequest, Partition: kafka.PartitionAny},
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
			var kafkaMessage data.KafkaResponseMessageUUIDData
			if err := json.Unmarshal(msg.Value, &kafkaMessage); err != nil {
				log.Println(err)
				continue
			}

			requestChannelInterface, ok := registerRequestChannels.Load(kafkaMessage.UUID)
			if !ok {
				log.Printf("Can't load channel %s from requestChannels", kafkaMessage.UUID)
				continue
			}
			requestChannel := requestChannelInterface.(chan string)
			requestChannel <- string(msg.Value)
		} else {
			// The client will automatically try to recover from all errors.
			log.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}

func registerResponseWatcher(id string, requestChannel chan string, errorChannel chan error) (*data.KafkaResponseMessage, []byte, error) {
	for {
		select {
		case err := <-errorChannel:
			return nil, nil, err
		case message := <-requestChannel:
			var kafkaMessage data.KafkaResponseMessage
			if err := json.Unmarshal([]byte(message), &kafkaMessage); err != nil {
				return nil, nil, err
			}
			registerRequestChannels.Delete(id)
			return &kafkaMessage, []byte(message), nil
		}
	}
}
