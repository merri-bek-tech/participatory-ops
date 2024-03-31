package main

import (
	"net/http"
	comps "parops/components"
	compCache "parops/components/component_cache"
	"parops/schemes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cache := compCache.NewComponentCache()

	e := echo.New()

	e.Pre(rewriteFrontEndPaths())
	e.Use(middleware.CORS())

	e.GET("/api", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/api/schemes", schemes.GetIndex)
	e.GET("/api/inbox", compCache.WithCache(comps.GetInbox, cache))

	e.Static("/", "/app/web")

	go comps.MonitorComponents(cache)

	e.Logger.Fatal(e.Start(":1323"))
}
