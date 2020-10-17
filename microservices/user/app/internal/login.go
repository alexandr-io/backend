package internal

import (
	"errors"
	"log"
	"net/http"

	"github.com/alexandr-io/berrors"

	"github.com/alexandr-io/backend/user/data"
	"github.com/alexandr-io/backend/user/database"
	"github.com/alexandr-io/backend/user/kafka/producers"

	"golang.org/x/crypto/bcrypt"
)

// Login is the internal logic function used to login a user.
func Login(key string, message data.KafkaUserLoginRequest) error {
	// Get the user from it's login
	user, err := database.GetUserByLogin(message.Data.Login)
	if err != nil {
		var badInput *data.BadInput
		if errors.As(err, &badInput) {
			return producers.SendBadRequestLoginMessage(key, badInput.JSONError)
		}
		return producers.SendInternalErrorLoginMessage(key, err.Error())
	}

	// Check the user's password
	if !comparePasswords(user.Password, []byte(message.Data.Password)) {
		return producers.SendBadRequestLoginMessage(key, berrors.BadInputJSONFromType("login", string(berrors.Login)))
	}

	return producers.SendSuccessLoginMessage(key, http.StatusOK, *user)
}

// comparePasswords compare a hashed password with a plain string password.
func comparePasswords(hashedPassword string, plainPassword []byte) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
