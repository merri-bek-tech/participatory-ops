package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/api", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Static("/", "/app/web")
	e.Logger.Fatal(e.Start(":1323"))
}
