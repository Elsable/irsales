package routes

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/nosnibor89/irsales/controllers"
)

// SetRoutes add routes(HttpHandler)
func SetRoutes(app *echo.Echo) {

	// UI/View routes
	app.Static("/", "ui/dist")

	// Api Routes
	api := app.Group("/api")

	api.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "This API is running")
	})

	companyController := controllers.NewCompanyController()

	api.POST("/company", companyController.SaveCompany)

	api.GET("/:name", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, "+c.Param("name"))
	})

	// f.GET("/users", actions.GetUsers)
}
