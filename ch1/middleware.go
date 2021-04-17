// ./ch1/middleware.go

// Middleware function
app.Use(func(c *fiber.Ctx) error {
	// print current data to console
	fmt.Println("Date:", time.Now())

	// passing the request to the next endpoint
	return c.Next()
})
