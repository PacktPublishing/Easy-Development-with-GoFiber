// ./ch1/static.js

app.use(
  "/static", // mount address
  express.static("public") // path to the file folder
);
