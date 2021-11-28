package main

import (
	"net/http"
	"path"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

// createRoute creates all the routes of the service.
func createRoute(app *fiber.App) {
	// Recover middleware in case of panic
	app.Use(recover.New())
	app.Use(limiter.New(limiter.Config{
		Max:				30,	
		Expiration:			15 * time.Second,
		LimitReached:		func(c *fiber.Ctx) error{
			return c.SendString("Rate limited, your IP is sending too many requests")
		},
	}))

	app.Get("/docs", mergeDocFiles, wrapDocHandler())
	app.Get("/merged/docs.yml", wrapFileServer())
	app.Get("/auth", mergeAuthFiles, wrapAuthDocHandler())
	app.Get("/merged/auth.yml", wrapFileServer())
	app.Get("/user", mergeUserFiles, wrapUserDocHandler())
	app.Get("/merged/user.yml", wrapFileServer())
	app.Get("/library", mergeLibraryFiles, wrapLibraryDocHandler())
	app.Get("/merged/library.yml", wrapFileServer())
	app.Get("/media", mergeMediaFiles, wrapMediaDocHandler())
	app.Get("/merged/media.yml", wrapFileServer())

	// Ping route used for testing that the service is up and running
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	// Custom 404 handler
	app.Use(func(c *fiber.Ctx) error {
		return fiber.ErrNotFound
	})
}

func fillRedocOpts(service string) middleware.RedocOpts {
	return middleware.RedocOpts{
		BasePath: "/",
		Path:     service,
		SpecURL:  path.Join("merged", service+".yml"),
		Title:    service + " documentation",
	}
}

func wrapFileServer() func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		fasthttpadaptor.NewFastHTTPHandler(http.FileServer(http.Dir("./")))(ctx.Context())
		return nil
	}
}

func wrapDocHandler() func(ctx *fiber.Ctx) error {
	swaggerHandler := middleware.Redoc(fillRedocOpts("docs"), nil)

	return func(ctx *fiber.Ctx) error {
		fasthttpadaptor.NewFastHTTPHandler(swaggerHandler)(ctx.Context())
		return nil
	}
}

func mergeDocFiles(ctx *fiber.Ctx) error {
	if err := mergeDocsFiles(ctx, "docs"); err != nil {
		return err
	}
	return ctx.Next()
}
