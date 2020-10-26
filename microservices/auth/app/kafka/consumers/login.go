package consumers

import (
	"fmt"
	"log"

	"github.com/alexandr-io/backend/auth/kafka/producer"
)

// consumeLoginResponseMessages consume all the messages coming to the `login-response` topic.
// Once a message is consumed, the UUID is extracted from the key to store the message to the correct loginRequestChannels channel.
func consumeLoginResponseMessages() {
	// Create new consumer
	consumer, err := newConsumer(loginResponse)
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

			requestChannelInterface, ok := producer.LoginRequestChannels.Load(string(msg.Key))
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
