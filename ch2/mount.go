// ./ch2/mount.go

package main

import "github.com/gofiber/fiber/v2"

func main() {
	// Create the first Fiber instance
	micro := fiber.New()

	// Create a new route for the first instance
	micro.Get("/doe", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	// Create the second Fiber instance
	app := fiber.New()

	// Create a new route for the second instance,
	// with included first instance's route
	app.Mount("/john", micro)

	// Start server on port 3000
	app.Listen(":3000")
}
