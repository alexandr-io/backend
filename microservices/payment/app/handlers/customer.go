package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/alexandr-io/backend/payment/data"
	"github.com/alexandr-io/backend/payment/database/customer"
	stripeCustomer "github.com/alexandr-io/backend/payment/stripe/customer"
)

// GetCustomerSubscription get the current subscription price of a user
func GetCustomerSubscription(ctx *fiber.Ctx) error {
	userID, err := userIDFromHeader(ctx)
	if err != nil {
		return err
	}

	var localCustomer *data.Customer
	localCustomer, err = customer.GetByUserID(userID)
	if err != nil {
		return err
	}

	price := stripeCustomer.GetCustomerSubscription(localCustomer.StripeID)

	if err = ctx.Status(fiber.StatusOK).Send([]byte(strconv.Itoa(int(price)))); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
