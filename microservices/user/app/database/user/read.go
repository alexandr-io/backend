package user

import (
	"context"
	"time"

	"github.com/alexandr-io/backend/user/data"
	"github.com/alexandr-io/backend/user/database"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// FromID get an user by it's given id.
func FromID(userID primitive.ObjectID) (*data.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := database.Instance.Db.Collection(database.CollectionUser)

	filter := bson.D{{Key: "_id", Value: userID}}
	var result data.User

	if err := collection.FindOne(ctx, filter).Decode(&result); err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusUnauthorized, err.Error())
	}
	return &result, nil
}

// FromEmail get an user by it's given email.
// In case of error, the proper error is set to the context and false is returned.
func FromEmail(login string) (*data.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := database.Instance.Db.Collection(database.CollectionUser)
	var object data.User

	// Get user by email
	emailFilter := bson.D{{Key: "email", Value: login}}
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

// FromLogin get an user by it's given login (username or email).
// In case of error, the proper error is set to the context and false is returned.
func FromLogin(login string) (*data.User, error) {
	if result, err := FromUsername(login); err == nil {
		return result, nil
	}
	if result, err := FromEmail(login); err == nil {
		return result, nil
	}

	return nil, data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "can't find user with login "+login)
}
