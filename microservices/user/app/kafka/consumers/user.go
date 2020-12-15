package consumers

import (
	"fmt"
	"log"

	"github.com/alexandr-io/backend/user/internal"
)

// consumeUserRequestMessages consume all the kafka message from the `user` topic.
// Once a message is consumed, it is sent to the user internal logic.
func consumeUserRequestMessages() {
	// Create new consumer
	consumer, err := newConsumer(userRequest)
	if err != nil {
		log.Println(err)
		return
	}
	defer consumer.Close()

	// Subscribe consumer to topic user
	if err := consumer.SubscribeTopics([]string{userRequest}, nil); err != nil {
		log.Println(err)
		return
	}

	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			fmt.Printf("[KAFKA]: Message on %s: %s:%s\n", msg.TopicPartition, string(msg.Key), string(msg.Value))

			// Send to logic
			_ = internal.User(string(msg.Key), string(msg.Value))
		} else {
			log.Printf("Consumer error: %s", err)
		}
	}
}
