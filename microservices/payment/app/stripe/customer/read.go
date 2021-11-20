package customer

import (
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/customer"
)

// GetCustomerSubscription get a customer subscription
func GetCustomerSubscription(stripeID string) int64 {
	params := &stripe.CustomerParams{}
	params.AddExpand("subscriptions")
	c, _ := customer.Get(stripeID, params)
	if len(c.Subscriptions.Data) != 0 {
		return c.Subscriptions.Data[0].Items.Data[0].Price.UnitAmount
	}
	return 0
}
