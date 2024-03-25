package main

import (
	"net/http"
	comps "parops/componentcache"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ComponentStatus struct {
	Uuid   string `json:"uuid"`
	Status string `json:"status"`
}

func main() {
	cache := make(map[string]comps.Component)

	e := echo.New()

	e.Pre(rewriteFrontEndPaths())
	e.Use(middleware.CORS())

	e.GET("/api", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/api/inbox", comps.WithCache(getInbox, cache))

	e.Static("/", "/app/web")

	go comps.PopulateComponentCache(cache)

	e.Logger.Fatal(e.Start(":1323"))
}

func getInbox(c echo.Context, cache comps.ComponentCache) error {
	components := make([]ComponentStatus, 0, len(cache))
	for _, cacheItem := range cache {
		component := ComponentStatus{
			Uuid:   cacheItem.Uuid,
			Status: comps.Status(cacheItem),
		}

		components = append(components, component)
	}

	return c.JSON(http.StatusOK, components)
}
