// ./ch1/hello_world.js

const express = require("express"); // add Express library
const app = express(); // create a new Express instance

// Create a new endpoint
app.get("/", (req, res) => {
  res.send("Hello World!"); // send text
});

// Start server on port 3000
app.listen(3000);
