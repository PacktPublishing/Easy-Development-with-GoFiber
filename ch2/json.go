// ./ch2/json.go

package main

import "github.com/gofiber/fiber/v2"

func main() {
	// Create a new Fiber instance
	app := fiber.New()

	// Create a new route with GET method
	app.Get("/json", func(c *fiber.Ctx) error {
		// Return response in JSON format
		return c.JSON(fiber.Map{
			"name": "John",
			"age":  33,
		})
	})

	// Start server on port 3000
	app.Listen(":3000")
}
