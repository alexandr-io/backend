package consumers

import (
	"fmt"
	"log"

	"github.com/alexandr-io/backend/auth/data"
	"github.com/alexandr-io/backend/auth/internal"
)

// consumeAuthRequestMessages consume all the kafka message from the `auth` topic.
// Once a message is consumed, it is sent to the auth internal logic.
func consumeAuthRequestMessages() {
	// Create new consumer
	consumer, err := newConsumer(authRequest)
	if err != nil {
		log.Println(err)
		return
	}
	defer consumer.Close()

	// Subscribe consumer to topic user
	if err := consumer.SubscribeTopics([]string{authRequest}, nil); err != nil {
		log.Println(err)
		return
	}

	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			fmt.Printf("[KAFKA]: Message on %s: %s:%s\n", msg.TopicPartition, string(msg.Key), string(msg.Value))
			messageData, err := data.GetAuthMessage(*msg)
			if err != nil {
				continue
			}
			// Send to logic
			err = internal.Auth(string(msg.Key), messageData)
			fmt.Printf("Error in Auth internal logic: %s\n", err.Error())
		} else {
			log.Printf("Topic: %s -> consumer error: %s", msg.TopicPartition, err)
		}
	}
}
