package main

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func rewriteFrontEndPaths() echo.MiddlewareFunc {
	paths := []string{"/api", "/assets"}

	return middleware.RewriteWithConfig(middleware.RewriteConfig{
		Skipper: func(c echo.Context) bool {
			for _, p := range paths {
				if strings.HasPrefix(c.Request().URL.Path, p) {
					return true
				}
			}
			if c.Request().Header.Get("Content-Type") == "application/json" {
				return true
			}
			return false
		},
		Rules: map[string]string{"^/*": "/"},
	})
}
