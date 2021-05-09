package invitation

import (
	"context"

	"github.com/alexandr-io/backend/auth/data"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Insert insert a new invitation into the database.
func Insert(collection *mongo.Collection, invitationData data.Invitation) (*data.Invitation, error) {
	insertedResult, err := collection.InsertOne(context.Background(), invitationData)
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	invitationData.ID = insertedResult.InsertedID.(primitive.ObjectID)
	return &invitationData, nil
}
