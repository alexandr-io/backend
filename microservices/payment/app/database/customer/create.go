package customer

import (
	"context"

	"github.com/alexandr-io/backend/payment/data"
	"github.com/alexandr-io/backend/payment/database"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Insert a customer in the database
func Insert(customer data.Customer) (*data.Customer, error) {
	result, err := database.CustomerCollection.InsertOne(context.Background(), customer)

	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	customer.ID = result.InsertedID.(primitive.ObjectID)
	return &customer, nil
}
