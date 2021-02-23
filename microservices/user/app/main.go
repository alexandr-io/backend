// license that can be found in the LICENSE file.

// User is the alexandrio microservice that handle all the users related features.
//
package main

import (
	"context"
	"log"

	"github.com/alexandr-io/backend/user/database"
	"github.com/alexandr-io/backend/user/grpc"
	"github.com/alexandr-io/backend/user/kafka/consumers"
	"github.com/alexandr-io/backend/user/kafka/producers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("User Service started")

	// MongoDB
	database.ConnectToMongo()
	defer database.Instance.Client.Disconnect(context.Background())
	database.InitCollections()

	// gRPC
	grpc.InitGRPC()
	defer grpc.CloseGRPC()

	consumers.StartConsumers()
	for producers.CreateTopics() != nil {
	}

	// Create a new fiber instance with custom config
	app := fiber.New(fiber.Config{
		// Override default error handler
		ErrorHandler: errorHandler,
	})
	createRoute(app)
	log.Fatal(app.Listen(":3000"))
}
