// ./ch2/stack.go

package main

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Create a new Fiber instance
	app := fiber.New()

	// Create new routes
	app.Get("/john/:age", handler)
	app.Post("/register", handler)

	// Print the router stack in JSON format
	data, _ := json.MarshalIndent(app.Stack(), "", "  ")
	fmt.Println(string(data))

	// Start server on port 3000
	app.Listen(":3000")
}
