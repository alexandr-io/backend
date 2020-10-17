package main

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

// createRoute creates all the routes of the service.
func createRoute(app *fiber.App) {

	app.Get("/docs", wrapDocHandler())
	app.Get("/swagger.yml", wrapFileServer())

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})
}

func wrapDocHandler() func(ctx *fiber.Ctx) error {
	options := middleware.RedocOpts{SpecURL: "/swagger.yml"}
	swaggerHandler := middleware.Redoc(options, nil)

	return func(ctx *fiber.Ctx) error {
		fasthttpadaptor.NewFastHTTPHandler(swaggerHandler)(ctx.Context())
		return nil
	}
}

func wrapFileServer() func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		fasthttpadaptor.NewFastHTTPHandler(http.FileServer(http.Dir("./")))(ctx.Context())
		return nil
	}
}
