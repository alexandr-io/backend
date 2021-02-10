package setters

import (
	"context"

	"github.com/alexandr-io/backend/auth/data"
	"github.com/alexandr-io/backend/auth/database/invitation"
	invitationGetters "github.com/alexandr-io/backend/auth/database/invitation/getters"
	"github.com/alexandr-io/backend/auth/database/mongo"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

// UpdateInvitation take a data.Invitation to update an invitation in database
func UpdateInvitation(invitationData data.Invitation) (data.Invitation, error) {
	// Update data
	invitationCollection := mongo.Instance.Db.Collection(invitation.Collection)
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
	updatedInvitation, err := invitationGetters.GetInvitationFromToken(invitationData.Token)
	if err != nil {
		return data.Invitation{}, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return *updatedInvitation, nil
}
