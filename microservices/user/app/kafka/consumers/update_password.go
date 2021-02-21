package consumers

import (
	"fmt"
	"log"

	"github.com/alexandr-io/backend/user/data"
	"github.com/alexandr-io/backend/user/internal"
)

// consumeUpdatePasswordRequestMessages consume all the kafka message from the `user.password.update` topic.
// Once a message is consumed, it is sent to the updatePassword internal logic.
func consumeUpdatePasswordRequestMessages() {
	// Create new consumer
	consumer, err := newConsumer(updatePasswordRequest)
	if err != nil {
		log.Println(err)
		return
	}
	defer consumer.Close()

	// Subscribe consumer to topic updatePassword
	if err := consumer.SubscribeTopics([]string{updatePasswordRequest}, nil); err != nil {
		log.Println(err)
		return
	}

	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			fmt.Printf("[KAFKA]: Message on %s: %s:%s\n", msg.TopicPartition, string(msg.Key), string(msg.Value))
			messageData, err := data.GetUserUpdatePasswordMessage(*msg)
			if err != nil {
				continue
			}
			// Send to logic
			_ = internal.UpdatePassword(string(msg.Key), messageData)
		} else {
			log.Printf("Consumer error: %s", err)
		}
	}
}
