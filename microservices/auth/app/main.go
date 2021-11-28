// license that can be found in the LICENSE file.

// User is the alexandrio microservice that handle all the users related features.
//
package main

import (
	"context"
	"log"
	"time"

	"github.com/alexandr-io/backend/auth/database"
	"github.com/alexandr-io/backend/auth/grpc"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Auth Service started")

	// MongoDB
	database.ConnectToMongo()
	defer database.Instance.Client.Disconnect(context.Background())
	database.InitCollections()

	// gRPC
	grpc.InitGRPC()
	defer grpc.CloseGRPC()

	// Create a new fiber instance with custom config
	app := fiber.New(fiber.Config{
		// Override default error handler
		ErrorHandler: errorHandler,
	})

	app.Use(limiter.New(limiter.Config{
		Max:				30,	
		Expiration:			15 * time.Second,
		LimitReached:		func(c *fiber.Ctx) error{
			return c.SendString("Rate limited, your IP is sending too many requests")
		},
	}))

	createRoute(app)

	log.Fatal(app.Listen(":3000"))
}
