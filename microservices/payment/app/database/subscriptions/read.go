package subscriptions

import (
	"context"

	"github.com/alexandr-io/backend/payment/data"
	"github.com/alexandr-io/backend/payment/database"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

// GetAll retrieve a list of every subscription available
func GetAll() ([]data.Subscription, error) {
	var subscriptions []data.Subscription

	cursor, err := database.SubscriptionsCollection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, "Cannot fetch data")
	}

	err = cursor.All(context.Background(), &subscriptions)
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, "Cannot read data")
	}

	return subscriptions, nil
}
