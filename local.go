package main

import (
	handler "public/api"
)

func main() {
	e := handler.Echo_route()
	e.Logger.Fatal(e.Start(":8080"))
}
