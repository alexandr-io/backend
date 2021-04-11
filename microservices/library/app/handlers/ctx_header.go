package handlers

import (
	"github.com/alexandr-io/backend/library/data"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func userIDFromHeader(ctx *fiber.Ctx) (primitive.ObjectID, error) {
	ID := string(ctx.Request().Header.Peek("ID"))
	userID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return userID, data.NewHTTPErrorInfo(fiber.StatusBadRequest, err.Error())
	}
	return userID, nil
}
