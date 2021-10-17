package customer

import (
	"github.com/alexandr-io/backend/payment/data"
	"github.com/stripe/stripe-go/v72"
	scustomer "github.com/stripe/stripe-go/v72/customer"
)

func Create(customer data.Customer) (*stripe.Customer, error) {
	params := &stripe.CustomerParams{
		Email:       stripe.String(customer.Email),
		Address: &stripe.AddressParams{
			City:       stripe.String(customer.Address.City),
			Country:    stripe.String(customer.Address.Country),
			Line1:      stripe.String(customer.Address.Line1),
			Line2:      stripe.String(customer.Address.Line2),
			PostalCode: stripe.String(customer.Address.PostalCode),
			State:      stripe.String(customer.Address.State),
		},
		Name:  stripe.String(customer.Username),
		Phone: stripe.String(customer.Phone),
	}

	params.AddMetadata("user_id", customer.UserID.Hex())

	result, err := scustomer.New(params)
	if err != nil {
		return nil, err
	}
	return result, nil
}