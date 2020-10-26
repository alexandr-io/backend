package consumers

import (
	"fmt"
	"log"

	"github.com/alexandr-io/backend/user/data"
	"github.com/alexandr-io/backend/user/internal"
)

// consumeLoginRequestMessages consume all the kafka message from the `login` topic.
// Once a message is consumed, it is sent to the login internal logic.
func consumeLoginRequestMessages() {
	// Create new consumer
	consumer, err := newConsumer(loginRequest)
	if err != nil {
		log.Println(err)
		return
	}
	defer consumer.Close()

	// Subscribe consumer to topic login
	if err := consumer.SubscribeTopics([]string{loginRequest}, nil); err != nil {
		log.Println(err)
		return
	}

	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			fmt.Printf("[KAFKA]: Message on %s: %s:%s\n", msg.TopicPartition, string(msg.Key), string(msg.Value))
			messageData, err := data.GetUserLoginMessage(*msg)
			if err != nil {
				continue
			}
			// Send to logic
			_ = internal.Login(string(msg.Key), messageData)
		} else {
			log.Printf("Topic: %s -> consumer error: %s", msg.TopicPartition, err)
		}
	}
}
