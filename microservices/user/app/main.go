// license that can be found in the LICENSE file.

// User is the alexandrio microservice that handle all the users related features.
//
package main

import (
	"context"
	"github.com/Alexandr-io/Backend/User/database"
	"log"

	"github.com/gofiber/fiber"
)

func main() {
	log.Println("User Service started")

	// MongoDB
	database.ConnectToMongo()
	defer database.Instance.Client.Disconnect(context.Background())
	database.InitCollections()

	// Fiber
	app := fiber.New()
	createRoute(app)
	app.Listen(3000)
}
