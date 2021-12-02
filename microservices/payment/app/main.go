// license that can be found in the LICENSE file.

// Media is the alexandrio microservice that handle all the media related features.
//
package main

import (
	"context"
	"log"
	"os"

	"github.com/alexandr-io/backend/payment/database"
	"github.com/alexandr-io/backend/payment/grpc"
	"github.com/alexandr-io/backend/payment/stripe"

	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Payment Service started")

	err := sentry.Init(sentry.ClientOptions{
		Dsn: "http://a3a4b09e28514398a147c7ac0c43eedf@95.217.135.159:9000/2",
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	// MongoDB
	database.ConnectToMongo()
	defer database.Instance.Client.Disconnect(context.Background())
	database.InitCollections()

	// gRPC
	grpc.InitGRPC()
	defer grpc.CloseGRPC()

	stripe.Setup(os.Getenv("STRIPE_API_KEY"))

	// Create a new fiber instance with custom config
	app := fiber.New(fiber.Config{
		// Override default error handler
		ErrorHandler: errorHandler,
	})
	createRoute(app)

	log.Fatal(app.Listen(":3000"))
}
