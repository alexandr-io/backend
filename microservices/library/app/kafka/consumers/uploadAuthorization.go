package consumers

import (
	"fmt"
	"log"

	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/internal"
	"github.com/alexandr-io/backend/library/kafka/producers"
	"github.com/gofiber/fiber/v2"
)

// consumeLibraryUploadAuthorizationMessage consume all the messages coming to the `library.upload.allowed` topic.
// Once a message is consumed, the UUID is extracted from the key to store the message to the correct LibraryUploadAuthorizationRequestChannels channel.
func consumeLibraryUploadAuthorizationMessage() {
	consumer, err := newConsumer(libraryUploadAuthorization)
	if err != nil {
		return
	}
	defer consumer.Close()

	// Subscribe consumer to topic 'library.upload.allowed'
	if err := consumer.SubscribeTopics([]string{libraryUploadAuthorization}, nil); err != nil {
		log.Println(err)
		return
	}

	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			fmt.Printf("[KAFKA]: Message on %s: %s:%s\n", msg.TopicPartition, string(msg.Key), string(msg.Value))
			messageData, err := data.GetUserLibraryUploadAuthorizationMessage(*msg)
			if err != nil {
				continue
			}

			isAllowed, err := internal.CanUserUploadOnLibrary(messageData)
			if err != nil {
				_ = producers.SendInternalErrorLibraryUploadAuthorizationMessage(string(msg.Key), err.Error())
				log.Println(err)
				continue
			}
			_ = producers.SendSuccessLibraryUploadAuthorizationMessage(string(msg.Key), fiber.StatusOK, isAllowed)
		} else {
			// The client will automatically try to recover from all errors.
			log.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}
