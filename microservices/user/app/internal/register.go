package internal

import (
	"encoding/json"
	"errors"
	"github.com/alexandr-io/backend/user/data"
	"github.com/alexandr-io/backend/user/database"
	"github.com/alexandr-io/backend/user/errorTypes"
	"github.com/alexandr-io/backend/user/kafka/producers"
	"github.com/alexandr-io/berrors"
	"log"
	"net/http"
)

func Register(messageID string, messageBytes []byte) {

	var userRegister data.UserRegister
	if err := json.Unmarshal(messageBytes, &userRegister); err != nil {
		log.Println(err)
		producers.SendKafkaMessageToProducer(messageID, berrors.KafkaErrorMessage{
			Code:    http.StatusInternalServerError,
			Content: []byte(err.Error()),
		})
		return
	}

	// Insert the new data to the collection
	insertedResult, err := database.InsertUserRegister(data.User{
		Username: userRegister.Username,
		Email:    userRegister.Email,
		Password: userRegister.Password,
	})
	if err != nil {
		var badInput *errorTypes.BadInput
		if errors.As(err, &badInput) {
			producers.SendKafkaMessageToProducer(messageID, berrors.KafkaErrorMessage{
				Code:    http.StatusBadRequest,
				Content: err.(*errorTypes.BadInput).JsonError,
			})
			return
		} else {
			producers.SendKafkaMessageToProducer(messageID, berrors.KafkaErrorMessage{
				Code:    http.StatusInternalServerError,
				Content: []byte(err.Error()),
			})
			return
		}
	}

	// Get the newly created user
	createdUser, ok := database.GetUserByID(insertedResult.InsertedID)
	if !ok {
		producers.SendKafkaMessageToProducer(messageID, berrors.KafkaErrorMessage{
			Code:    http.StatusInternalServerError,
			Content: []byte("internal server error after user insertion"),
		})
		return
	}

	createdUserJson, err := json.Marshal(createdUser)
	if err != nil {
		producers.SendKafkaMessageToProducer(messageID, berrors.KafkaErrorMessage{
			Code:    http.StatusInternalServerError,
			Content: []byte(err.Error()),
		})
		return
	}

	producers.SendKafkaMessageToProducer(messageID, berrors.KafkaErrorMessage{
		Code:    http.StatusCreated,
		Content: createdUserJson,
	})
}
