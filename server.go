package main

import (
	"github.com/labstack/echo"
	"github.com/nosnibor89/irsales/config"
)

func main() {
	e := echo.New()
	f := config.Framework{e}

	config.New(f)

	f.Logger.Fatal(e.Start(":3000"))
}
