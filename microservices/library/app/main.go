// license that can be found in the LICENSE file.

// Library is the alexandrio microservice that handle all the libraries related features.
//
package main

import (
	"context"
	"log"

	"github.com/alexandr-io/backend/library/database"
	"github.com/alexandr-io/backend/library/grpc"
	"github.com/alexandr-io/backend/library/kafka/consumers"
	"github.com/alexandr-io/backend/library/kafka/producers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Library Service started")

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
	createRoute(app)

	consumers.StartConsumers()
	for producers.CreateTopics() != nil {
	}

	log.Fatal(app.Listen(":3000"))
}
