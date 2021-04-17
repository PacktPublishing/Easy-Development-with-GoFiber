// ./ch1/routing.go

// Route for Post method
app.Post("/", func(c *fiber.Ctx) error {
	// function that stores a new data
})

// Route for PUT method
app.Put("/", func(c *fiber.Ctx) error {
	// function that replaces the existing data
})

// Route for PATH method
app.Path("/", func(c *fiber.Ctx) error {
	// function that replaces part of the existing data
})

// Route for DELETE method
app.Delete("/", func(c *fiber.Ctx) error {
	// function that deletes the data
})