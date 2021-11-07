package internal

import (
	"github.com/alexandr-io/backend/payment/data"
	"github.com/alexandr-io/backend/payment/database/customer"
	scustomer "github.com/alexandr-io/backend/payment/stripe/customer"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateStripeCustomerForUser(user *data.User, customerData data.Customer) (*data.Customer, error) {
	userID, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusBadRequest, err.Error())
	}
	customerData.UserID = userID
	customerData.Email = user.Email
	customerData.Username = user.Username

	stripeCustomer, err := scustomer.Create(customerData)
	if err != nil {
		return nil, err
	}

	customerData.StripeID = stripeCustomer.ID
	localCustomer, err := customer.Insert(customerData)
	if err != nil {
		return nil, err
	}
	return localCustomer, nil
}
