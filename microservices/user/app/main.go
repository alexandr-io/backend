// license that can be found in the LICENSE file.

// User is the alexandrio microservice that handle all the users related features.
//
package main

import (
	"context"
	"log"

	"github.com/alexandr-io/backend/user/database"
	"github.com/alexandr-io/backend/user/grpc"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
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

	// Templates
	templates := html.New("./templates", ".html")

	// Create a new fiber instance with custom config
	app := fiber.New(fiber.Config{
		// Override default error handler
		ErrorHandler: errorHandler,
		Views:        templates,
	})
	createRoute(app)
	log.Fatal(app.Listen(":3000"))
}
