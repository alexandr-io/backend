package invitation

import (
	"context"

	"github.com/alexandr-io/backend/auth/data"
	"github.com/alexandr-io/backend/auth/database"
	"github.com/alexandr-io/backend/auth/database/mongo"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

// Update take a data.Invitation to update an invitation in database
func Update(invitationData data.Invitation) (data.Invitation, error) {
	// Update data
	invitationCollection := mongo.Instance.Db.Collection(database.CollectionInvitation)
	if _, err := invitationCollection.UpdateOne(
		context.Background(),
		bson.D{
			{"token", invitationData.Token},
		},
		bson.D{
			{"$set", invitationData},
		},
	); err != nil {
		return data.Invitation{}, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	// Retrieve updated data
	updatedInvitation, err := GetFromToken(invitationData.Token)
	if err != nil {
		return data.Invitation{}, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return *updatedInvitation, nil
}
