package config

import (
	"net/http"

	"github.com/labstack/echo"
)

// SetRoutes add routes(HttpHandler)
func SetRoutes(app Framework) {
	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	app.GET("/test/:name", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, "+c.Param("name"))
	})

	// f.GET("/users", actions.GetUsers)
}
