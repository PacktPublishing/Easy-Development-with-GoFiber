// ./ch2/routing.go

package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New() // create a new Fiber instance

	// Simple route:
	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Route with named parameters:
	app.Get("/user/:author/books/:title", func(c *fiber.Ctx) error {
		str := fmt.Sprintf("%s, %s", c.Params("author"), c.Params("title"))
		return c.SendString(str)
	})

	// Route with optional named parameter:
	app.Get("/user/:name?", func(c *fiber.Ctx) error {
		return c.SendString(c.Params("name"))
	})

	// Plus sign for NOT optional greedy parameter:
	app.Get("/user/+", func(c *fiber.Ctx) error {
		return c.SendString(c.Params("+"))
	})

	// Wildcard sign for optional greedy parameter:
	app.Get("/user/*", func(c *fiber.Ctx) error {
		return c.SendString(c.Params("*"))
	})

	// Complex route example:
	app.Get("/flights/:from-:to/time::at", func(c *fiber.Ctx) error {
		str := fmt.Sprintf(
			"%s-%s at %s", c.Params("from"), c.Params("to"), c.Params("at"),
		)
		return c.SendString(str)
	})

	// Escaped route:
	app.Get("/resource/key\\:value", func(c *fiber.Ctx) error {
		return c.SendString("escaped key:value")
	})

	// Start server on port 3000
	app.Listen(":3000")
}
