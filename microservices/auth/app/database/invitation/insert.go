package invitation

import (
	"context"
	"time"

	"github.com/alexandr-io/backend/auth/data"
	"github.com/alexandr-io/backend/auth/database"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Insert insert a new invitation into the database.
func Insert(invitationData data.Invitation) (*data.Invitation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	invitationCollection := database.Instance.Db.Collection(database.CollectionInvitation)

	insertedResult, err := invitationCollection.InsertOne(ctx, data.Invitation{
		Token:  invitationData.Token,
		Used:   nil,
		UserID: nil,
	})
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	invitationData.ID = insertedResult.InsertedID.(primitive.ObjectID).Hex()
	return &invitationData, nil
}
