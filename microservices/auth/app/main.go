// license that can be found in the LICENSE file.

// User is the alexandrio microservice that handle all the users related features.
//
package main

import (
	"log"

	"github.com/alexandr-io/backend/auth/kafka"

	"github.com/gofiber/fiber/v2"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Auth Service started")

	app := fiber.New()
	createRoute(app)

	kafka.StartConsumers()

	log.Fatal(app.Listen(":3000"))
}
