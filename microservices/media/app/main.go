// license that can be found in the LICENSE file.

// Media is the alexandrio microservice that handle all the media related features.
//
package main

import (
	"context"
	"log"

	"github.com/alexandr-io/backend/media/database"
	"github.com/alexandr-io/backend/media/grpc"

	"github.com/gofiber/fiber/v2"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Media Service started")

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
		BodyLimit:    20 * 1024 * 1024, // 20MB max upload
	})
	createRoute(app)

	log.Fatal(app.Listen(":3000"))
}
