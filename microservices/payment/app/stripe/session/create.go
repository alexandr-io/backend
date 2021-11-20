package session

import (
	"github.com/alexandr-io/backend/payment/data"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/checkout/session"
)

// Create a stripe customer session
func Create(sessionData data.Session) (*stripe.CheckoutSession, error) {
	params := &stripe.CheckoutSessionParams{
		SuccessURL: stripe.String(sessionData.SuccessURL),
		CancelURL:  stripe.String(sessionData.CancelURL),
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				Price:    stripe.String(sessionData.PriceID),
				Quantity: stripe.Int64(1),
			},
		},
		Mode:     stripe.String(string(stripe.CheckoutSessionModeSubscription)),
		Customer: stripe.String(sessionData.CustomerID),
	}

	result, err := session.New(params)
	if err != nil {
		return nil, err
	}
	return result, nil
}
