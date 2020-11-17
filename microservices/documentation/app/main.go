package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/miracl/conflate"
)

func init() {
	// define the unmarshalers for the given file extensions
	conflate.Unmarshallers = conflate.UnmarshallerMap{
		".yaml": {conflate.YAMLUnmarshal},
		".yml":  {conflate.YAMLUnmarshal},
	}
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Auth Service started")

	// Create a new fiber instance with custom config
	app := fiber.New()
	createRoute(app)

	log.Fatal(app.Listen(":2999"))
}
