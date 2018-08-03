package app

import (
	"github.com/labstack/echo"
	"github.com/nosnibor89/irsales/repository"
	"github.com/nosnibor89/irsales/routes"
)

// Run does all the heavy lifting for the app. Set routes, connect to DB, set logging, set middlewares, etc
func Run() {
	app := echo.New()
	routes.SetRoutes(app)

	session := repository.Connect()
	defer session.Close()

	app.Logger.Fatal(app.Start(":3000"))
}
