package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func main() {

	// Create fiber app
	app :=  fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	fmt.Println("Starting app")

	app.Listen(":3000")
}