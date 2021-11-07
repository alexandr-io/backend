package handlers

import (
	"log"

	"github.com/alexandr-io/backend/payment/data"
	"github.com/alexandr-io/backend/payment/database/customer"
	"github.com/alexandr-io/backend/payment/internal"

	"github.com/gofiber/fiber/v2"
)

func Subscribe(ctx *fiber.Ctx) error {
	user := userFromHeader(ctx)
	userID, err := userIDFromHeader(ctx)
	if err != nil {
		return err
	}

	var Customer data.Customer
	err = ParseBodyJSON(ctx, &Customer)
	if err != nil {
		return err
	}

	var localCustomer *data.Customer
	localCustomer, err = customer.GetByUserID(userID)
	if err != nil {
		if e, ok := err.(*fiber.Error); ok {
			// Override status code if fiber.Error type
			if e.Code == fiber.StatusNotFound {

				localCustomer, err = internal.CreateStripeCustomerForUser(user, Customer)
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

	log.Println(localCustomer)
	// TODO: CREATE LINK TO PAY ON STRIPE
	//
	// Should use localCustomer
	return nil
}
