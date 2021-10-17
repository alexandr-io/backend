package customer

import (
	"context"
	"github.com/alexandr-io/backend/payment/data"
	"github.com/alexandr-io/backend/payment/database"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Update(customerData data.Customer) (*data.Customer, error) {
	if err := database.CustomerCollection.FindOneAndUpdate(
		context.Background(),
		bson.D{
			{"_id", customerData.ID},
		},
		bson.D{{"$set", customerData}},
		options.FindOneAndUpdate().SetReturnDocument(1),
	).Decode(&customerData); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, data.NewHTTPErrorInfo(fiber.StatusNotFound, "Customer not found.")
		}
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return &customerData, nil
}