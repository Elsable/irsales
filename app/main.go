package app

import (
	"github.com/labstack/echo"
	"github.com/nosnibor89/irsales/repository"
	"github.com/nosnibor89/irsales/routes"
)

// Run does all the heavy lifting for the app. Set routes, connect to DB, set logging, set middlewares, etc
func Run() {
	app := echo.New()

	// Let's set DB connection first
	session := repository.Connect()
	defer session.Close()

	routes.SetRoutes(app)
	app.Logger.Fatal(app.Start(":3000"))
}
