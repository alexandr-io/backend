package customer

import (
	"context"

	"github.com/alexandr-io/backend/payment/data"
	"github.com/alexandr-io/backend/payment/database"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetByUserID get a stripe user from its alexandrio ID
func GetByUserID(userID primitive.ObjectID) (*data.Customer, error) {
	var Customer data.Customer

	customerFilter := bson.D{{"user_id", userID}}
	if err := database.CustomerCollection.FindOne(context.Background(), customerFilter).Decode(&Customer); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, data.NewHTTPErrorInfo(fiber.StatusNotFound, "Customer not found.")
		}
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	return &Customer, nil
}
