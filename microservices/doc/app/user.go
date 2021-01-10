package main

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

const user = "user"

func wrapUserDocHandler() func(ctx *fiber.Ctx) error {
	swaggerHandler := middleware.Redoc(fillRedocOpts(user), nil)

	return func(ctx *fiber.Ctx) error {
		fasthttpadaptor.NewFastHTTPHandler(swaggerHandler)(ctx.Context())
		return nil
	}
}

func mergeUserFiles(ctx *fiber.Ctx) error {
	if err := mergeDocsFiles(ctx, user); err != nil {
		return err
	}
	return ctx.Next()
}
