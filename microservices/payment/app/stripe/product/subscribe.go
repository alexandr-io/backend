package product

import (
	"github.com/alexandr-io/backend/payment/data"
	"github.com/gofiber/fiber/v2"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/sub"
)

func Subscribe(customer *data.Customer) (*stripe.Subscription, error) {
	params := &stripe.SubscriptionParams{
		Customer: stripe.String(customer.StripeID),
		Items: []*stripe.SubscriptionItemsParams{
			{
				Price: stripe.String("price_1Ja52IBybdEu8beGqKHy8c9L"),
			},
		},
	}
	s, err := sub.New(params)
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return s, nil

}