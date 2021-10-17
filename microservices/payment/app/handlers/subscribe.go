package handlers

import (
	"github.com/alexandr-io/backend/payment/data"
	"github.com/alexandr-io/backend/payment/database/customer"
	"github.com/alexandr-io/backend/payment/internal"
	customer2 "github.com/alexandr-io/backend/payment/stripe/customer"
	"github.com/alexandr-io/backend/payment/stripe/paymentmethod"
	"github.com/alexandr-io/backend/payment/stripe/product"
	"github.com/gofiber/fiber/v2"
)

func Subscribe(ctx *fiber.Ctx) error {
	user := userFromHeader(ctx)
	userID, err := userIDFromHeader(ctx)
	if err != nil {
		return err
	}

	var Subscription data.Subscribe
	err = ParseBodyJSON(ctx, &Subscription)
	if err != nil {
		return err
	}

	var localCustomer *data.Customer
	localCustomer, err = customer.GetByUserID(userID)
	if err != nil {
		if e, ok := err.(*fiber.Error); ok {
			// Override status code if fiber.Error type
			if e.Code == fiber.StatusNotFound {

				localCustomer, err = internal.CreateStripeCustomerForUser(user, Subscription.Customer)
				if err != nil {
					return err
				}
			} else {
				return err
			}
		} else {
			return err
		}
	}

	if Subscription.CreditCard.ID == "" {
		resultCard, err := paymentmethod.Create(Subscription.CreditCard)
		if err != nil {
			return err
		}
		Subscription.CreditCard.ID = resultCard.ID

		err = paymentmethod.Attach(localCustomer.StripeID, resultCard.ID)
		if err != nil {
			return err
		}
	} else {
		_, err := paymentmethod.GetFromID(Subscription.CreditCard.ID)
		if err != nil {
			return err
		}
	}

	_, err = customer2.UpdateDefaultPaymentMethod(localCustomer.StripeID, Subscription.CreditCard.ID)
	if err != nil {
		return err
	}

	_, err = product.Subscribe(localCustomer)
	if err != nil {
		return err
	}
	return nil
}
