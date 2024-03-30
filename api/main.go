package main

import (
	"net/http"
	comps "parops/components"
	compCache "parops/components/component_cache"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"parops.libs/msg"
)

type ComponentStatus struct {
	Uuid    string                `json:"uuid"`
	Status  string                `json:"status"`
	Details *msg.ComponentDetails `json:"details"`
}

func main() {
	cache := compCache.NewComponentCache()

	e := echo.New()

	e.Pre(rewriteFrontEndPaths())
	e.Use(middleware.CORS())

	e.GET("/api", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/api/inbox", compCache.WithCache(getInbox, cache))

	e.Static("/", "/app/web")

	go comps.MonitorComponents(cache)

	e.Logger.Fatal(e.Start(":1323"))
}

func getInbox(c echo.Context, cache *compCache.ComponentCache) error {
	components := cache.ItemList()
	statuses := make([]ComponentStatus, 0, cache.ItemCount())

	for _, cacheItem := range components {
		statuses = append(statuses, ComponentStatus{
			Uuid:    cacheItem.Uuid,
			Status:  compCache.Status(cacheItem),
			Details: cacheItem.Details,
		})
	}

	return c.JSON(http.StatusOK, statuses)
}
