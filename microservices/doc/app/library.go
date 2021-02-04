package main

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

const library = "library"

func wrapLibraryDocHandler() func(ctx *fiber.Ctx) error {
	swaggerHandler := middleware.Redoc(fillRedocOpts(library), nil)

	return func(ctx *fiber.Ctx) error {
		fasthttpadaptor.NewFastHTTPHandler(swaggerHandler)(ctx.Context())
		return nil
	}
}

func mergeLibraryFiles(ctx *fiber.Ctx) error {
	if err := mergeDocsFiles(ctx, library); err != nil {
		return err
	}
	return ctx.Next()
}
