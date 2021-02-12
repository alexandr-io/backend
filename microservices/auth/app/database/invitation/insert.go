package invitation

import (
	"context"
	"time"

	"github.com/alexandr-io/backend/auth/data"
	"github.com/alexandr-io/backend/auth/database"
	mongo2 "github.com/alexandr-io/backend/auth/database/mongo"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

// Insert insert a new invitation into the database.
func Insert(invitationData data.Invitation) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	invitationCollection := mongo2.Instance.Db.Collection(database.CollectionInvitation)

	insertedResult, err := invitationCollection.InsertOne(ctx, data.Invitation{
		Token:  invitationData.Token,
		Used:   nil,
		UserID: nil,
	})
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return insertedResult, nil
}
