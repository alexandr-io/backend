package main

import (
	"net/http"

	"github.com/alexandr-io/backend/user/handlers"
	userMiddleware "github.com/alexandr-io/backend/user/middleware"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

// createRoute creates all the routes of the service.
func createRoute(app *fiber.App) {
	// Recover middleware in case of panic
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		TimeFormat: "2 Jan 15:04:05 MST",
		TimeZone:   "Europe/Paris",
		Next: func(ctx *fiber.Ctx) bool {
			if string(ctx.Request().RequestURI()) == "/dashboard" ||
				string(ctx.Request().RequestURI()) == "/docs" ||
				string(ctx.Request().RequestURI()) == "/swagger.yml" {
				return true
			}
			return false
		},
	}))
	app.Get("/dashboard", monitor.New())
	app.Get("/docs", wrapDocHandler())
	app.Get("/swagger.yml", wrapFileServer())

	app.Get("/me", userMiddleware.Protected(), handlers.Me)

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	// Custom 404 handler
	app.Use(func(c *fiber.Ctx) error {
		return fiber.ErrNotFound
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
