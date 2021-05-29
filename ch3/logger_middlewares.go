// ./ch3/security_middlewares.go

package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Create a new Fiber instance
	app := fiber.New()

	// Define file to logs
	file, err := os.OpenFile("./my_logs.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()

	// Set config for logger
	loggerConfig := logger.Config{
		Format:     "${pid} ${status} - ${method} ${path}\n", // set logs format
		TimeFormat: "02-Jan-2006",                            // set time format
		TimeZone:   "America/New_York",                       // set time zone
		Output:     file,                                     // add file to save output
	}

	// Use middlewares for each route
	app.Use(
		logger.New(loggerConfig), // add Logger middleware with config
	)

	// Create a new endpoint
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hey, logger!") // send text
	})

	// Start server on port 3000
	app.Listen(":3000")
}
