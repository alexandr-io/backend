package handlers

import (
	"github.com/alexandr-io/backend/payment/data"
	"github.com/alexandr-io/backend/payment/internal"

	"github.com/gofiber/fiber/v2"
)

func ListSubscriptions(ctx *fiber.Ctx) error {

	result, err := internal.GetMergedSubscriptions()
	if err != nil {
		return err
	}

	if err = ctx.Status(fiber.StatusOK).JSON(result); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
