package product

import (
	"github.com/alexandr-io/backend/common/typeconv"
	"github.com/alexandr-io/backend/payment/data"
	"github.com/gofiber/fiber/v2"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/price"
	"github.com/stripe/stripe-go/v72/product"
)

func GetFromID(stripeID string) (*stripe.Product, error) {
	result, err := product.Get(stripeID, nil)
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusNotFound, "Product not found.")
	}
	return result, nil
}

func GetProductPrices(stripeID string) []stripe.Price {
	var prices []stripe.Price

	params := &stripe.PriceListParams{
		Active:       typeconv.BoolPtr(true),
		Product:      &stripeID,
	}
	result := price.List(params)
	for result.Next() {
		prices = append(prices, *result.Price())
	}
	return prices
}