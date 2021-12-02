package invitation

import (
	"context"

	"github.com/alexandr-io/backend/auth/data"
	"github.com/alexandr-io/backend/auth/database"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Insert inserts a new invitation into the database.
func Insert(invitationData data.Invitation) (*data.Invitation, error) {
	insertedResult, err := database.InvitationCollection.InsertOne(context.Background(), invitationData)
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	invitationData.ID = insertedResult.InsertedID.(primitive.ObjectID)
	return &invitationData, nil
}
