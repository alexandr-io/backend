package invitation

import (
	"context"
	"github.com/alexandr-io/backend/auth/database"

	"github.com/alexandr-io/backend/auth/data"
	mongo2 "github.com/alexandr-io/backend/auth/database/mongo"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

// Delete delete an invitation corresponding to the given invitation token
func Delete(token string) error {
	invitationCollection := mongo2.Instance.Db.Collection(database.CollectionInvitation)

	result, err := invitationCollection.DeleteOne(
		context.Background(),
		bson.D{
			{Key: "token", Value: token},
		},
	)
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	if result.DeletedCount == 0 {
		return data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "can't find invitation to delete")
	}
	return nil
}
