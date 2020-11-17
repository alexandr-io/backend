package consumers

import (
	"fmt"
	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/internal"
	"log"
)

// consumeLibrariesCreationMessages consume all the kafka message from the `login` topic.
// Once a message is consumed, it is sent to the login internal logic.
func consumeLibrariesCreationMessages() {
	// Create new consumer
	consumer, err := newConsumer(librariesRequest)
	if err != nil {
		log.Println(err)
		return
	}
	defer consumer.Close()

	// Subscribe consumer to topic login
	if err := consumer.SubscribeTopics([]string{librariesRequest}, nil); err != nil {
		log.Println(err)
		return
	}

	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			fmt.Printf("[KAFKA]: Message on %s: %s:%s\n", msg.TopicPartition, string(msg.Key), string(msg.Value))
			messageData, err := data.GetUserLibrariesCreationMessage(*msg)
			if err != nil {
				continue
			}
			// Send to logic
			_ = internal.CreateLibraries(string(msg.Key), messageData)
		} else {
			log.Printf("Topic: %s -> consumer error: %s", msg.TopicPartition, err)
		}
	}
}
