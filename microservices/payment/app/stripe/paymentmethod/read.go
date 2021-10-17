package paymentmethod

import (
	"github.com/alexandr-io/backend/payment/data"
	"github.com/gofiber/fiber/v2"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/paymentmethod"
)

func GetCustomerCards(customerID string) []stripe.PaymentMethod {

	var cards []stripe.PaymentMethod
	params := &stripe.PaymentMethodListParams{
		Customer: stripe.String(customerID),
		Type: stripe.String("card"),
	}
	i := paymentmethod.List(params)
	for i.Next() {
		cards = append(cards, *i.PaymentMethod())
	}
	return cards
}

func GetFromID(paymentMethodID string) (*stripe.PaymentMethod, error) {
	pm, err := paymentmethod.Get(
		paymentMethodID,
		nil,
	)
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusNotFound, "Card not found.")
	}
	return pm, nil
}