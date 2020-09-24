package consumers

import (
	"fmt"
	"github.com/alexandr-io/backend/user/data"
	"github.com/alexandr-io/backend/user/internal"
	"log"
)

func consumeRegisterRequestMessages() {
	// Create new consumer
	consumer, err := newConsumer()
	if err != nil {
		log.Println(err)
		return
	}
	defer consumer.Close()

	// Subscribe consumer to topic register-
	if err := consumer.SubscribeTopics([]string{registerRequest}, nil); err != nil {
		log.Println(err)
		return
	}

	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
			messageData, err := data.GetMessageFromBytes(msg.Value)
			if err != nil {
				continue
			}
			// Send to logic
			internal.Register(messageData.UUID, messageData.Data)
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}
