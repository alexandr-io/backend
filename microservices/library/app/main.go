package main

import (
	"context"
	"log"

	"github.com/alexandr-io/backend/library/kafka/consumers"

	"github.com/alexandr-io/backend/library/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Library Service started")

	// MongoDB
	database.ConnectToMongo()
	defer database.Instance.Client.Disconnect(context.Background())
	database.InitCollections()

	consumers.StartConsumers()

	// Fiber
	app := fiber.New(fiber.Config{
		// Override default error handler
		ErrorHandler: errorHandler,
	})

	createRoute(app)
	log.Fatal(app.Listen(":3000"))
}
