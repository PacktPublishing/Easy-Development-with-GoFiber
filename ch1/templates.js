// ./ch1/templates.js

// Initialize Pug template engine for rendering
app.set("view engine", "pug");

// Initialize templates folder
app.set("views", "./views");

// Create a new endpoint
app.get("/", (req, res) => {
  // rendering the "index" template with content passing
  res.render("index", {
    title: "Hey!",
    message: "This is the index template.",
  });
});
