package main

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

const auth = "auth"

func wrapAuthDocHandler() func(ctx *fiber.Ctx) error {
	swaggerHandler := middleware.Redoc(fillRedocOpts(auth), nil)

	return func(ctx *fiber.Ctx) error {
		fasthttpadaptor.NewFastHTTPHandler(swaggerHandler)(ctx.Context())
		return nil
	}
}

func mergeAuthFiles(ctx *fiber.Ctx) error {
	if err := mergeDocsFiles(ctx, auth); err != nil {
		return err
	}
	return ctx.Next()
}
