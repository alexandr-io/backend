package user

import (
	"context"
	"time"

	"github.com/alexandr-io/backend/user/data"
	"github.com/alexandr-io/backend/user/database"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// FromUsername get an user by it's given username.
// In case of error, the proper error is set to the context and false is returned.
func FromUsername(username string) (*data.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := database.Instance.Db.Collection(database.CollectionUser)
	var object data.User

	// Get user by username
	emailFilter := bson.D{{Key: "username", Value: username}}
	// Return a login error if the user is not found
	if err := collection.FindOne(ctx, emailFilter).Decode(&object); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, data.NewHTTPErrorInfo(fiber.StatusNotFound, "User not found")
		}
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	// Return the email user object
	return &object, nil
}
