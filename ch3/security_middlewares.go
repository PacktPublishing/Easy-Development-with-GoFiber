// ./ch3/security_middlewares.go

package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/helmet/v2"
	"github.com/gofiber/utils"
)

func main() {
	// Create a new Fiber instance
	app := fiber.New()

	// CSRF configuration
	csrfConfig := csrf.Config{
		KeyLookup:      "header:X-Csrf-Token", // string in the form of '<source>:<key>' that is used to extract token from the request
		CookieName:     "my_csrf_",            // name of the session cookie
		CookieSameSite: "Strict",              // indicates if CSRF cookie is requested by SameSite
		Expiration:     1 * time.Hour,         // expiration is the duration before CSRF token will expire
		KeyGenerator:   utils.UUID,            // creates a new CSRF token
	}

	// Use middlewares for each route
	app.Use(
		helmet.New(),         // add Helmet middleware
		csrf.New(csrfConfig), // add CSRF middleware with config
	)

	// Create a new endpoint
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("It's secure!") // send text
	})

	// Start server on port 3000
	app.Listen(":3000")
}
