// license that can be found in the LICENSE file.

// User is the alexandrio microservice that handle all the users related features.
//
package main

import (
	"context"
	"log"

	"github.com/Alexandr-io/Backend/User/database"

	"github.com/gofiber/fiber"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
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
