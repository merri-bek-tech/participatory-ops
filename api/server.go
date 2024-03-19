package main

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type UnknownComponent struct {
	Uuid   string `json:"uuid"`
	Status string `json:"status"`
}

func main() {
	e := echo.New()

	e.Pre(rewriteExcept([]string{"/api", "/assets"}, map[string]string{"^/*": "/"}))

	e.GET("/api", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/api/inbox", getInbox)

	e.Static("/", "/app/web")
	e.Logger.Fatal(e.Start(":1323"))
}

func getInbox(c echo.Context) error {
	components := []UnknownComponent{
		{Uuid: "f08b7172-36d8-447f-85e1-41403d2730c8", Status: "online"},
	}
	return c.JSON(http.StatusOK, components)
}

func rewriteExcept(paths []string, rules map[string]string) echo.MiddlewareFunc {
	return middleware.RewriteWithConfig(middleware.RewriteConfig{
		Skipper: func(c echo.Context) bool {
			for _, p := range paths {
				if strings.HasPrefix(c.Request().URL.Path, p) {
					return true
				}
			}
			return false
		},
		Rules: rules,
	})
}
