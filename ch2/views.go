// ./ch2/views.go

package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
	// Initialize a standard Go html template engine
	engine := html.New("./views", ".html")

	// Create a new Fiber template with template engine
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		// Render a template named 'index.html' with content
		return c.Render("index", fiber.Map{
			"Title":       "Hello, World!",
			"Description": "This is a template.",
		})
	})

	// Start server on port 3000
	app.Listen(":3000")
}
