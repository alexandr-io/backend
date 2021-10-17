package customer

import (
	"github.com/alexandr-io/backend/payment/data"
	"github.com/gofiber/fiber/v2"
	"github.com/stripe/stripe-go/v72"
	scustomer "github.com/stripe/stripe-go/v72/customer"

)

func Retrieve(customer data.Customer) (*stripe.Customer, error) {
	result, err := scustomer.Get(customer.StripeID, nil)
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusNotFound, "Customer not found")
	}
	return result, nil
}