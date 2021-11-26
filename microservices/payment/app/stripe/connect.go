package stripe

import (
	"errors"

	"github.com/stripe/stripe-go/v72"
)

// Setup stripe connection
func Setup(APIKey string) error {
	if APIKey == "" {
		return errors.New("invalid Stripe's API key")
	}
	stripe.Key = APIKey
	return nil
}
