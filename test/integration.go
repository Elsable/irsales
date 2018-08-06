// Package test initializes the App and resources for testing
package test

import (
	"github.com/labstack/echo"
)

// TestApp is a placeholder for all the integration tests and/or tests that uses the app object.
// For example: Controllers, routes, context
var TestApp *echo.Echo

func init() {
	TestApp = echo.New()
}
