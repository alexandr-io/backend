package kafka

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/alexandr-io/backend/auth/data"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

var loginRequestChannels sync.Map

// LoginRequestHandler is the entry point of a new login message to the user MS using kafka.
// The function produce a new message to kafka,
// create a channel for the answer,
// call a watcher to wait for the proper answer from the login-response topic,
// interpret the answer (possible errors or success) and return and error with the proper http code
// In case of success, a data.User is returned containing the username and email of the user.
func LoginRequestHandler(user data.UserLogin) (*data.KafkaUser, error) {
	// Generate UUID
	id := uuid.New()
	// Create a channel for the request
	requestChannel := make(chan string)
	// Create a channel to manage error in goroutines
	errorChannel := make(chan error)
	// Store request channel to the channel map
	loginRequestChannels.Store(id.String(), requestChannel)
	// Produce the message to kafka
	go produceLoginMessage(id.String(), user, errorChannel)
	// Watch for a response in the request channel
	kafkaMessage, rawMessage, err := loginResponseWatcher(id.String(), requestChannel, errorChannel)
	if err != nil {
		return nil, err
	}

	// handle error
	if err := handleError(*kafkaMessage, rawMessage); err != nil {
		return nil, err
	}

	// handle success
	if kafkaMessage.Data.Code == http.StatusOK {
		return data.UnmarshalUserResponse(rawMessage)
	}

	// If http code contained in the kafka message is not handled
	return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, fmt.Sprintf("unmanaged code: %d", kafkaMessage.Data.Code))
}

// produceLoginMessage produce a login message to the `login` topic.
// The message sent is a JSON of the data.KafkaUserLoginMessage struct.
func produceLoginMessage(id string, user data.UserLogin, errorChannel chan error) {
	message, err := data.CreateLoginMessage(user)
	if err != nil {
		errorChannel <- data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
		return
	}

	// Create a new producer
	producer, err := newProducer()
	if err != nil {
		errorChannel <- data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
		return
	}
	defer producer.Close()

	// Delivery report handler for produced messages
	go produceMessageReport(producer)

	// Produce message to topic (asynchronously)
	if err := producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &loginRequest, Partition: kafka.PartitionAny},
		Key:            []byte(id),
		Value:          message,
	}, nil); err != nil {
		errorChannel <- data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
		return
	}

	// Wait for message deliveries before shutting down
	producer.Flush(int((15 * time.Microsecond).Microseconds()))
	return
}

// consumeLoginResponseMessages consume all the messages coming to the `login-response` topic.
// Once a message is consumed, the UUID is extracted from the key to store the message to the correct loginRequestChannels channel.
func consumeLoginResponseMessages() {
	// Create new consumer
	consumer, err := newConsumer()
	if err != nil {
		return
	}
	defer consumer.Close()

	// Subscribe consumer to topic login-response
	if err := consumer.SubscribeTopics([]string{loginResponse}, nil); err != nil {
		log.Println(err)
		return
	}

	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			fmt.Printf("[KAFKA]: Message on %s: %s:%s\n", msg.TopicPartition, string(msg.Key), string(msg.Value))

			requestChannelInterface, ok := loginRequestChannels.Load(string(msg.Key))
			if !ok {
				log.Printf("Can't load channel %s from requestChannels", msg.Key)
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

// loginResponseWatcher is waiting for an update in the given channel. The message will be set in the channel by
// consumeLoginResponseMessages once the user MS has answered to the request.
// The channel is then deleted from the map and the kafka message is returned.
func loginResponseWatcher(id string, requestChannel chan string, errorChannel chan error) (*data.KafkaResponseMessage, []byte, error) {
	timeout := time.After(5 * time.Second)
	for {
		select {
		case <-timeout:
			// In case of time out, we delete the channel and return an error
			loginRequestChannels.Delete(id)
			return nil, nil, data.NewHTTPErrorInfo(fiber.StatusGatewayTimeout, "Kafka login response timed out")
		case err := <-errorChannel:
			return nil, nil, err
		case message := <-requestChannel:
			var kafkaMessage data.KafkaResponseMessage
			if err := json.Unmarshal([]byte(message), &kafkaMessage); err != nil {
				return nil, nil, err
			}
			loginRequestChannels.Delete(id)
			return &kafkaMessage, []byte(message), nil
		}
	}
}
