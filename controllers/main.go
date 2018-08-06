package controllers

import (
	"fmt"

	"github.com/labstack/echo"
)

// Alive return a map with dummy message for a controller's action
func Alive(c echo.Context) map[string]string {

	aliveIn := fmt.Sprintf("you're in %s  %s", c.Request().Method, c.Path())

	return map[string]string{"message": aliveIn}
}
