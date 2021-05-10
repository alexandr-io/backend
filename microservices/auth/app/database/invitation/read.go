package invitation

import (
	"context"

	"github.com/alexandr-io/backend/auth/data"
	"github.com/alexandr-io/backend/auth/database"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetFromToken get an invitation by it's given token.
func GetFromToken(token string) (*data.Invitation, error) {
	filter := bson.D{{Key: "token", Value: token}}
	object := &data.Invitation{}

	if err := database.InvitationCollection.FindOne(context.Background(), filter).Decode(object); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, data.NewHTTPErrorInfo(fiber.StatusNotFound, "Invitation not found")
		}
		return nil, data.NewHTTPErrorInfo(fiber.StatusUnauthorized, err.Error())
	}
	return object, nil
}
