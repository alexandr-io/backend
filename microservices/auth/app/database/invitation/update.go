package invitation

import (
	"context"

	"github.com/alexandr-io/backend/auth/data"
	"github.com/alexandr-io/backend/auth/database"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Update take a data.Invitation to update an invitation in database
func Update(invitationData data.Invitation) (*data.Invitation, error) {
	if err := database.InvitationCollection.FindOneAndUpdate(
		context.Background(),
		bson.D{
			{"token", invitationData.Token},
		},
		bson.D{
			{"$set", invitationData},
		},
		options.FindOneAndUpdate().SetReturnDocument(1),
	).Decode(&invitationData); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, data.NewHTTPErrorInfo(fiber.StatusNotFound, "Invitation not found.")
		}
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return &invitationData, nil
}
