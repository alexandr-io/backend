package main

import "github.com/gofiber/fiber"

func main() {
	app := fiber.New()

	app.Get("/ping", func(c *fiber.Ctx) {
		c.Send("pong")
	})

	app.Listen(3000)
}
