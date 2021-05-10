package user

import (
	"context"

	"github.com/alexandr-io/backend/user/data"
	"github.com/alexandr-io/backend/user/database"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Update take a data.User and an ID to update a user in database
func Update(id primitive.ObjectID, user data.User) (*data.User, error) {
	// Update data
	userCollection := database.Instance.Db.Collection(database.CollectionUser)
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", user}}

	if err := userCollection.FindOneAndUpdate(context.Background(), filter, update, options.FindOneAndUpdate().SetReturnDocument(1)).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, data.NewHTTPErrorInfo(fiber.StatusNotFound, "User not found")
		}
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return &user, nil
}
