package producers

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// Topic define the structure used to define a kafka's topic
type Topic struct {
	// Name of the topic created in the kafka broker
	Name string
	// RetentionMS is the delay in milliseconds of the topic's documents time to live
	// set to 0 to use kafka default value (~7 days)
	// set to -1 to keep the documents forever
	RetentionMS int
	// NumPartitions in the topic.
	NumPartitions int
	// ReplicationFactor of the topic
	ReplicationFactor int
}

// ToTopicSpecification transform a Topic to a kafka.TopicSpecification
func (topic *Topic) ToTopicSpecification() kafka.TopicSpecification {
	return kafka.TopicSpecification{
		Topic:             topic.Name,
		NumPartitions:     topic.NumPartitions,
		ReplicationFactor: topic.ReplicationFactor,
		Config:            map[string]string{"retention.ms": strconv.Itoa(topic.RetentionMS)},
	}
}

var (
	// OLD: registerResponse = "register-response"
	registerResponse = Topic{
		Name:              "user.register.response",
		RetentionMS:       1000 * 15,
		NumPartitions:     1,
		ReplicationFactor: 1,
	}
	// OLD: loginResponse    = "login-response"
	loginResponse = Topic{
		Name:              "user.login.response",
		RetentionMS:       1000 * 15,
		NumPartitions:     1,
		ReplicationFactor: 1,
	}
	// OLD: userResponse     = "user-response"
	userResponse = Topic{
		Name:              "user.retrieve.response",
		RetentionMS:       900000, // 15min
		NumPartitions:     1,
		ReplicationFactor: 1,
	}
	// OLD: authRequest      = "auth"
	authRequest = Topic{
		Name:              "auth.token",
		RetentionMS:       1000 * 15,
		NumPartitions:     1,
		ReplicationFactor: 1,
	}

	updatePasswordResponse = Topic{
		Name:              "user.password.update.response",
		RetentionMS:       1000 * 15,
		NumPartitions:     1,
		ReplicationFactor: 1,
	}
)

// CreateTopics wait for the kafka broker to be running and create the topics
func CreateTopics() error {
	client, err := kafka.NewAdminClient(&kafka.ConfigMap{"bootstrap.servers": os.Getenv("KAFKA_URL")})
	if err != nil {
		log.Println("Error while creating topics: " + err.Error())
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	durationBeforeTimeout, err := time.ParseDuration("60s")
	if err != nil {
		log.Println("Error while creating topics: " + err.Error())
		return err
	}
	results, err := client.CreateTopics(
		ctx,
		[]kafka.TopicSpecification{
			registerResponse.ToTopicSpecification(),
			loginResponse.ToTopicSpecification(),
			userResponse.ToTopicSpecification(),
			authRequest.ToTopicSpecification(),
			updatePasswordResponse.ToTopicSpecification(),
		},
		kafka.SetAdminOperationTimeout(durationBeforeTimeout))
	if err != nil {
		log.Println("Error while creating topics: " + err.Error())
		return err
	}

	for _, result := range results {
		if result.Error.Code() != kafka.ErrNoError {
			log.Println(result.String())
		} else {
			log.Println("Topic " + result.String() + " Created.")
		}
	}
	return nil
}
