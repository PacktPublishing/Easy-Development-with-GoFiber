// ./ch2/body_parser.go

package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// Field names should start with an uppercase letter
type Person struct {
	Name  string `json:"name" xml:"name" form:"name"`
	Email string `json:"email" xml:"email" form:"email"`
}

func main() {
	// Create a new Fiber instance
	app := fiber.New()

	// Create a new route with POST method
	app.Post("/create", func(c *fiber.Ctx) error {
		// Define a new Person struct
		person := new(Person)

		// Binds the request body to the Person struct
		if err := c.BodyParser(person); err != nil {
			return err
		}

		// Print data
		fmt.Println(person.Name, person.Email)

		return nil
	})

	// Start server on port 3000
	app.Listen(":3000")
}
