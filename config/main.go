package config

import "github.com/labstack/echo"

// Framework is a wrapper for the underliying framework
type Framework struct {
	*echo.Echo
}

// New sets the frameworks app object and start adding routes,middlewarers,etc
func New(app Framework) {
	SetRoutes(app)
	SetDB()
}
