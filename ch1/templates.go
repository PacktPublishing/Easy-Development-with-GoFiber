// ./ch1/templates.go

// Initialize Pug template engine in ./views folder
engine := pug.New("./views", ".pug")

// Create a new Fiber instance
app := fiber.New(fiber.Config{
	Views: engine, // set template engine for rendering
})

// Create a new endpoint
app.Get("/", func(c *fiber.Ctx) error {
	// rendering the "index" template with content passing
	return c.Render("index", fiber.Map{
		"Title":   "Hey!",
		"Message": "This is the index template.",
	})
})
