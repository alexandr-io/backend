// license that can be found in the LICENSE file.

// Media is the alexandrio microservice that handle all the media related features.
//
package main

import (
	"context"
	"log"

	"github.com/alexandr-io/backend/media/database"
	consumers "github.com/alexandr-io/backend/media/kafka/comsumers"
	"github.com/alexandr-io/backend/media/kafka/producers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Media Service started")

	// MongoDB
	database.ConnectToMongo()
	defer database.Instance.Client.Disconnect(context.Background())
	database.InitCollections()

	consumers.StartConsumers()
	err := producers.CreateTopics()
	for err != nil {
		err = producers.CreateTopics()
	}

	// Create a new fiber instance with custom config
	app := fiber.New(fiber.Config{
		// Override default error handler
		ErrorHandler: errorHandler,
	})
	createRoute(app)

	log.Fatal(app.Listen(":3000"))
}
