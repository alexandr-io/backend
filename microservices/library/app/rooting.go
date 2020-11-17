package main

import (
	"net/http"

	"github.com/alexandr-io/backend/library/handlers"
	userMiddleware "github.com/alexandr-io/backend/library/middleware"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

func createRoute(app *fiber.App) {
	// Recover middleware in case of panic
	app.Use(recover.New())
	//15 Oct 09:36:56 CEST | 504 |   523ms |      172.19.0.1 | POST    | /register        | {"message":"Kafka register response timed out","file":"/app/kafka/register.go","line":143,"content-type":"text/html; charset=utf-8","custom-message":""}
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

	app.Post("/library", userMiddleware.Protected(), handlers.LibraryCreation)
	app.Get("/library", userMiddleware.Protected(), handlers.LibraryRetrieve)
	app.Delete("/library", userMiddleware.Protected(), handlers.LibraryDelete)

	app.Get("/libraries", userMiddleware.Protected(), handlers.LibrariesRetrieve)

	app.Get("/docs", wrapDocHandler())
	app.Get("/swagger.yml", wrapFileServer())

	// Ping route used for testing that the service is up and running
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
