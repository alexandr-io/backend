package main

import (
	"net/http"

	"github.com/Alexandr-io/Backend/User/handlers"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gofiber/fiber"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

// createRoute creates all the routes of the service.
func createRoute(app *fiber.App) {

	app.Post("/register", handlers.Register)

	app.Get("/docs", wrapDocHandler())
	app.Get("/swagger.yml", wrapFileServer())

	app.Get("/ping", func(c *fiber.Ctx) {
		c.Send("pong")
	})
}

func wrapDocHandler() func(ctx *fiber.Ctx) {
	options := middleware.RedocOpts{SpecURL: "/swagger.yml"}
	swaggerHandler := middleware.Redoc(options, nil)

	return func(ctx *fiber.Ctx) {
		fasthttpadaptor.NewFastHTTPHandler(swaggerHandler)(ctx.Fasthttp)
	}
}

func wrapFileServer() func(ctx *fiber.Ctx) {
	return func(ctx *fiber.Ctx) {
		fasthttpadaptor.NewFastHTTPHandler(http.FileServer(http.Dir("./")))(ctx.Fasthttp)
	}
}
