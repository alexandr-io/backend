// license that can be found in the LICENSE file.

// User is the alexandrio microservice that handle all the users related features.
//
package main

import (
	"log"

	"github.com/alexandr-io/backend/auth/kafka/consumers"
	"github.com/alexandr-io/backend/auth/kafka/producers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Auth Service started")

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
