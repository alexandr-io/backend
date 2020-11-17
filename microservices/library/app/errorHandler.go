package main

import (
	"fmt"

	"github.com/alexandr-io/backend/library/data"

	"github.com/gofiber/fiber/v2"
)

// handle the http error of fiber and log the errors
func errorHandler(ctx *fiber.Ctx, err error) error {
	// Default 500 statuscode
	code := fiber.StatusInternalServerError

	if e, ok := err.(*fiber.Error); ok {
		// Override status code if fiber.Error type
		code = e.Code
	}

	// Log the error
	errorInfo, infoErr := data.ErrorInfoUnmarshal(err.Error())
	if infoErr == nil {
		fmt.Printf("[HTTP ERROR]: %d %s -> [%s:%d] %s\n", code, ctx.Request().RequestURI(), errorInfo.File, errorInfo.Line, errorInfo.Message)
	} else {
		fmt.Printf("[HTTP ERROR]: %d %s -> %s\n", code, ctx.Request().RequestURI(), err.Error())
	}

	// Set Content-Type: text/plain; charset=utf-8
	ctx.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
	if errorInfo != nil && errorInfo.ContentType != fiber.MIMETextPlainCharsetUTF8 {
		ctx.Set(fiber.HeaderContentType, errorInfo.ContentType)
	}

	// Return default message for 500+ error
	if code >= fiber.StatusInternalServerError || code == fiber.StatusUnauthorized {
		return ctx.SendStatus(code)
	}
	// Send custom message if set or message contained in the errorInfo
	if errorInfo != nil && errorInfo.CustomMessage != "" {
		return ctx.Status(code).SendString(errorInfo.CustomMessage)
	} else if errorInfo != nil {
		return ctx.Status(code).SendString(errorInfo.Message)
	}
	// Send error message
	return ctx.Status(code).SendString(err.Error())
}
