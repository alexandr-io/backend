package consumers

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/alexandr-io/backend/mail/data"
	"github.com/alexandr-io/backend/mail/internal"
)

// ConsumeMailRequestMessages consume all the messages coming to the `mail.new` topic.
// Once a message is consumed, the content is sent to the internal logic who will create and send to mail.
func ConsumeMailRequestMessages() {
	// Create new consumer
	consumer, err := newConsumer(mailRequest.Name)
	if err != nil {
		return
	}
	defer consumer.Close()

	// Subscribe consumer to topic mail.new
	if err := consumer.SubscribeTopics([]string{mailRequest.Name}, nil); err != nil {
		log.Println(err)
		return
	}

	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			fmt.Printf("[KAFKA]: Message on %s: %s:%s\n", msg.TopicPartition, string(msg.Key), string(msg.Value))

			var mailData data.KafkaMail
			if err := json.Unmarshal(msg.Value, &mailData); err != nil {
				log.Println(err)
			}
			internal.CreateMailFromMessage(mailData)
		} else {
			// The client will automatically try to recover from all errors.
			log.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}
