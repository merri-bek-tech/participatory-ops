package main

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// rewrite middleware with skipper
	e.Pre(middleware.RewriteWithConfig(middleware.RewriteConfig{
		Skipper: func(c echo.Context) bool {
			return strings.HasPrefix(c.Request().URL.Path, "/api") || strings.HasPrefix(c.Request().URL.Path, "/assets")
		},
		Rules: map[string]string{
			"^/*": "/",
		},
	}))

	e.GET("/api", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Static("/", "/app/web")
	e.Logger.Fatal(e.Start(":1323"))
}
