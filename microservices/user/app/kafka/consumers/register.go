package consumers

import (
	"fmt"
	"log"

	"github.com/alexandr-io/backend/user/data"
	"github.com/alexandr-io/backend/user/internal"
)

// consumeRegisterRequestMessages consume all the kafka message from the `register` topic.
// Once a message is consumed, it is sent to the register internal logic.
func consumeRegisterRequestMessages() {
	// Create new consumer
	consumer, err := newConsumer()
	if err != nil {
		log.Println(err)
		return
	}
	defer consumer.Close()

	// Subscribe consumer to topic register
	if err := consumer.SubscribeTopics([]string{registerRequest}, nil); err != nil {
		log.Println(err)
		return
	}

	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			fmt.Printf("[KAFKA]: Message on %s: %s:%s\n", msg.TopicPartition, string(msg.Key), string(msg.Value))
			messageData, err := data.GetUserRegisterMessage(*msg)
			if err != nil {
				continue
			}
			// Send to logic
			_ = internal.Register(string(msg.Key), messageData)
		} else {
			log.Printf("Topic: %s -> consumer error: %s", msg.TopicPartition, err)
		}
	}
}
