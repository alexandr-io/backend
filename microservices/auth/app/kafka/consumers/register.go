package consumers

import (
	"fmt"
	"log"

	"github.com/alexandr-io/backend/auth/kafka/producers"
)

// consumeRegisterResponseMessages consume all the messages coming to the `register-response` topic.
// Once a message is consumed, the UUID is extracted from the key to store the message to the correct registerRequestChannels channel.
func consumeRegisterResponseMessages() {
	// Create new consumer
	consumer, err := newConsumer(registerResponse)
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
			fmt.Printf("[KAFKA]: Message on %s: %s:%s\n", msg.TopicPartition, string(msg.Key), string(msg.Value))

			requestChannelInterface, ok := producers.RegisterRequestChannels.Load(string(msg.Key))
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
