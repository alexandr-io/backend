package internal

import (
	"github.com/alexandr-io/backend/payment/data"
	"github.com/alexandr-io/backend/payment/database/subscriptions"
	"github.com/alexandr-io/backend/payment/stripe/product"
)

func GetMergedSubscriptions() (*[]data.Subscription, error) {
	var result []data.Subscription

	localProducts, err := subscriptions.GetAll()
	if err != nil {
		return nil, err
	}

	for _, localProductIter := range localProducts {
		stripeProduct, err := product.GetFromID(localProductIter.StripeID)
		if err != nil {
			return nil, err
		}

		prices := product.GetProductPrices(stripeProduct.ID)
		var pricesFormat []data.Price
		for _, price := range prices {
			pricesFormat = append(pricesFormat, data.Price{
				ID: price.ID,
				Currency:  price.Currency,
				Recurring: data.Recurring{
					Interval:      price.Recurring.Interval,
					IntervalCount: price.Recurring.IntervalCount,
				},
				Price:     price.UnitAmount,
			})
		}

		result = append(result, data.Subscription{
			ID: localProductIter.ID,
			Name:        stripeProduct.Name,
			Description: stripeProduct.Description,
			Prices: pricesFormat,
		})
	}
	return &result, nil
}