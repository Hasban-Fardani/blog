package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

//	func main() {
//		e := echo.New()
//		e.Static("/", "public")
//		e.Static("/icons", "static/icons")
//		e.Logger.Fatal(e.Start(":1323"))
//	}
func Handler(w http.ResponseWriter, r *http.Request) {
	e := echo.New()
	e.Static("/", "public")
	e.Static("/icons", "static/icons")

	e.ServeHTTP(w, r)
}
