// ./ch1/middleware.go

// Middleware function
app.use(function (req, res, next) {
  // print current data to console
  console.log("Date:", Date.now());

  // passing the request to the next endpoint
  next();
});
