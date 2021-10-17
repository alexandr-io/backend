package handlers

import (
	"github.com/alexandr-io/backend/payment/database/customer"
	"github.com/alexandr-io/backend/payment/stripe/paymentmethod"
	"github.com/gofiber/fiber/v2"
	"log"
)

func ListCards(ctx *fiber.Ctx) error {
	userID, err := userIDFromHeader(ctx)
	if err != nil {
		return err
	}

	localCustomer, err := customer.GetByUserID(userID)
	if err != nil {
		return err
	}
	cards := paymentmethod.GetCustomerCards(localCustomer.StripeID)
	log.Println(cards)
	return nil
}