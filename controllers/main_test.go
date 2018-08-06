package controllers_test

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"

	"github.com/nosnibor89/irsales/controllers"

	"github.com/nosnibor89/irsales/test"
)

func TestAlive(t *testing.T) {
	app := test.TestApp

	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	ctx := app.NewContext(req, rec)

	expected := fmt.Sprintf("you're in %s  %s", ctx.Request().Method, ctx.Path())

	result := controllers.Alive(ctx)

	if result["message"] != expected {
		t.Errorf("Failed Alive: expeted %s but got %s", result["message"], expected)
	}
}
