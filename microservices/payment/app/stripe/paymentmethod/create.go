package paymentmethod

import (
	"github.com/alexandr-io/backend/payment/data"
	"github.com/gofiber/fiber/v2"
	"github.com/stripe/stripe-go/v72"
	spaymentmethod "github.com/stripe/stripe-go/v72/paymentmethod"
)

func Create(card data.CreditCard) (*stripe.PaymentMethod, error) {
	params := &stripe.PaymentMethodParams{
		Card: &stripe.PaymentMethodCardParams{
			Number: stripe.String(card.Number),
			ExpMonth: stripe.String(card.ExpMonth),
			ExpYear: stripe.String(card.ExpYear),
			CVC: stripe.String(card.CVC),
		},
		Type: stripe.String("card"),
	}
	result, err := spaymentmethod.New(params)
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return result, nil
}

func Attach(customerID string, paymentMethodID string) error {
	params := &stripe.PaymentMethodAttachParams{
		Customer: stripe.String(customerID),
	}
	_, err := spaymentmethod.Attach(
		paymentMethodID,
		params,
	)
	return err
}