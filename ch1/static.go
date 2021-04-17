// ./ch1/static.go

app.Static(
	"/static",  // mount address
	"./public", // path to the file folder
)