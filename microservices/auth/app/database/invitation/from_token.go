package invitation

import (
	"context"
	"github.com/alexandr-io/backend/auth/database"
	"time"

	"github.com/alexandr-io/backend/auth/data"
	"github.com/alexandr-io/backend/auth/database/mongo"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

// GetFromToken get an invitation by it's given token.
func GetFromToken(token string) (*data.Invitation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := mongo.Instance.Db.Collection(database.CollectionInvitation)
	filter := bson.D{{Key: "token", Value: token}}
	object := &data.Invitation{}

	if err := collection.FindOne(ctx, filter).Decode(object); err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusUnauthorized, err.Error())
	}
	return object, nil
}
