package consumers

import (
	"fmt"
	"log"

	"github.com/alexandr-io/backend/user/data"
	"github.com/alexandr-io/backend/user/internal"
)

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
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
			messageData, err := data.GetUserRegisterMessage(*msg)
			if err != nil {
				continue
			}
			// Send to logic
			_ = internal.Register(messageData)
		} else {
			log.Printf("Topic: %s -> consumer error: %s", msg.TopicPartition, err)
		}
	}
}
