package consumers

import (
	"fmt"
	"log"

	"github.com/alexandr-io/backend/media/kafka/producers"
)

// consumeAuthResponseMessages consume all the messages coming to the `auth-response` topic.
// Once a message is consumed, the UUID is extracted from the key to store the message to the correct AuthRequestChannels channel.
func consumeAuthResponseMessages() {
	// Create new consumer
	consumer, err := newConsumer(authGroup)
	if err != nil {
		return
	}
	defer consumer.Close()

	// Subscribe consumer to topic Auth-response
	if err := consumer.SubscribeTopics([]string{authResponse}, nil); err != nil {
		log.Println(err)
		return
	}

	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			fmt.Printf("[KAFKA]: Message on %s: %s:%s\n", msg.TopicPartition, string(msg.Key), string(msg.Value))

			requestChannelInterface, ok := producers.AuthRequestChannels.Load(string(msg.Key))
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