package handlers

import (
	"github.com/alexandr-io/backend/payment/data"
	"github.com/alexandr-io/backend/payment/database/customer"
	"github.com/alexandr-io/backend/payment/internal"
	"github.com/alexandr-io/backend/payment/stripe/session"

	"github.com/gofiber/fiber/v2"
)

// Subscribe get a URL for a payment instance for a given produce and price
func Subscribe(ctx *fiber.Ctx) error {
	user := userFromHeader(ctx)
	userID, err := userIDFromHeader(ctx)
	if err != nil {
		return err
	}

	var sessionData data.Session
	err = ParseBodyJSON(ctx, &sessionData)
	if err != nil {
		return err
	}

	var customerData = data.Customer{
		UserID:   userID,
		Email:    user.Email,
		Username: user.Username,
	}

	var localCustomer *data.Customer
	localCustomer, err = customer.GetByUserID(userID)
	if err != nil {
		if e, ok := err.(*fiber.Error); ok {
			// Override status code if fiber.Error type
			if e.Code == fiber.StatusNotFound {

				localCustomer, err = internal.CreateStripeCustomerForUser(user, customerData)
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

	sessionData.CustomerID = localCustomer.StripeID
	newSession, err := session.Create(sessionData)
	if err != nil {
		return err
	}

	if err = ctx.Status(fiber.StatusOK).Send([]byte(newSession.URL)); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
