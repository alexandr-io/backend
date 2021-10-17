package data

import (
	"github.com/stripe/stripe-go/v72"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Subscribe struct {
	CreditCard   CreditCard         `json:"credit_card,omitempty" validate:"required"`
	Customer     Customer           `json:"customer,omitempty" validate:"required"`
}

type SubscriptionNew struct {
	ID    primitive.ObjectID `json:"id"`
	Price string `json:"price_id"`
}

type Subscription struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string             `json:"name,required"`
	Description string             `json:"description,required"`
	StripeID    string             `json:"-" bson:"stripe_id"`
	Prices      []Price            `json:"prices"`
}

type Price struct {
	ID        string          `json:"id"`
	Currency  stripe.Currency `json:"currency"`
	Recurring Recurring       `json:"recurring"`
	Price     int64           `json:"price"`
}

type Recurring struct {
	Interval      stripe.PriceRecurringInterval `json:"interval"`
	IntervalCount int64                         `json:"interval_count"`
}
