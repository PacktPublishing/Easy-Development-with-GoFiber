// ./ch2/query.go

package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Create a new Fiber instance
	app := fiber.New()

	// Create a new route with query params
	app.Get("/user", func(c *fiber.Ctx) error {
		// Print name from query param
		fmt.Println(c.Query("name"))

		return nil
	})

	// Start server on port 3000
	app.Listen(":3000")
}
