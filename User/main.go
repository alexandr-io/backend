// license that can be found in the LICENSE file.

// User is the alexandrio microservice that handle all the users related features.
//
package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber"
)

func main() {
	log.Println("User Service started")

	// MongoDB
	connectToMongo()
	defer instanceMongo.Client.Disconnect(context.Background())
	initCollections()

	// Fiber
	app := fiber.New()
	createRoute(app)
	app.Listen(3000)
}
