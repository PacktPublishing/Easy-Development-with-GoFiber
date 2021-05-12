// ./ch2/group.go

package main

import "github.com/gofiber/fiber/v2"

func main() {
	// Create a new Fiber instance
	app := fiber.New()

	// Create a new route group '/api'
	api := app.Group("/api", handler)

	// Create a new route for API v1
	v1 := api.Group("/v1", handler)
	v1.Get("/list", handler)

	// Create a new route for API v1
	v2 := api.Group("/v2", handler)
	v2.Get("/list", handler)

	// Start server on port 3000
	app.Listen(":3000")
}
