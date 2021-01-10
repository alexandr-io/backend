package main

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

const media = "media"

func wrapMediaDocHandler() func(ctx *fiber.Ctx) error {
	swaggerHandler := middleware.Redoc(fillRedocOpts(media), nil)

	return func(ctx *fiber.Ctx) error {
		fasthttpadaptor.NewFastHTTPHandler(swaggerHandler)(ctx.Context())
		return nil
	}
}

func mergeMediaFiles(ctx *fiber.Ctx) error {
	if err := mergeDocsFiles(ctx, media); err != nil {
		return err
	}
	return ctx.Next()
}
