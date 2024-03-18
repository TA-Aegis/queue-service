package main

import "github.com/gofiber/fiber/v3"

func main() {
	app := fiber.New()

	queue := app.Group("/bc/queue")

	queue.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	queue.Get("/ping", func(c fiber.Ctx) error {
		return c.SendString("pong!")
	})

	app.Listen(":8080")
}
