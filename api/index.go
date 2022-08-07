package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Echo_route() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/", "public")
	return e
}

func Handler(w http.ResponseWriter, r *http.Request) {
	e := Echo_route()
	e.ServeHTTP(w, r)
}
