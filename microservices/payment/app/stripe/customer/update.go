package customer

import (
	"github.com/alexandr-io/backend/payment/data"
	"github.com/gofiber/fiber/v2"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/customer"
)

func UpdateDefaultPaymentMethod(customerID string, paymentMethodID string) (*stripe.Customer, error) {
	params := &stripe.CustomerParams{
		InvoiceSettings:     &stripe.CustomerInvoiceSettingsParams{
			DefaultPaymentMethod: &paymentMethodID,
		},
	}
	c, err := customer.Update(
		customerID,
		params,
	)
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return c, nil
}