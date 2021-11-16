package data

import "go.mongodb.org/mongo-driver/bson/primitive"

// Address of a customer
type Address struct {
	City       string `json:"city"`
	Country    string `json:"country"`
	Line1      string `json:"line_1"`
	Line2      string `json:"line_2"`
	PostalCode string `json:"postal_code"`
	State      string `json:"state"`
}

// Customer to send to stripe
type Customer struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	UserID       primitive.ObjectID `json:"-" bson:"user_id"`
	Email        string             `json:"-"`
	Username     string             `json:"-"`
	Phone        string             `json:"phone"`
	Address      Address            `json:"address"`
	StripeID     string             `json:"-" bson:"stripe_id"`
	Subscription SubscriptionNew    `json:"subscription"`
}
