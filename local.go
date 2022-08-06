// for running web on local for testing
package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/", "public")
	e.GET("/login", func(c echo.Context) error {
		return c.HTML(200, "Login page")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
