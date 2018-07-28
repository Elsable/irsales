package config

import (
	"net/http"

	"github.com/labstack/echo"
)

// SetRoutes add routes(HttpHandler)
func SetRoutes(app Framework) {

	// UI/View routes
	app.Static("/", "ui/dist")

	// Api Routes
	api := app.Group("/api")

	api.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "This API is running")
	})

	api.GET("/:name", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, "+c.Param("name"))
	})

	// f.GET("/users", actions.GetUsers)
}
