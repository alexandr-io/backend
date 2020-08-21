package main

import (
	"github.com/alexandr-io/backend_errors"
	"github.com/gofiber/fiber"
	"net/http"
)

type UserRegister struct {
	Email           string `json:"email" validate:"required,email"`
	Username        string `json:"username" validate:"required"`
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" validate:"required"`
}

func register(ctx *fiber.Ctx) {
	ctx.Set("Content-Type", "application/json")

	userRegister := new(UserRegister)
	if ok := backend_errors.ParseBodyJSON(ctx, userRegister); !ok {
		return
	}

	if userRegister.Password != userRegister.ConfirmPassword {
		errorData := backend_errors.BadInputJSON("confirm_password", "passwords does not match")
		ctx.Status(http.StatusBadRequest).SendBytes(errorData)
		return
	}

	if err := ctx.Status(200).JSON("JWT"); err != nil {
		ctx.SendStatus(http.StatusInternalServerError)
	}
}
