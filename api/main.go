package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ComponentStatus struct {
	Uuid   string `json:"uuid"`
	Status string `json:"status"`
}

func main() {
	cache := make(map[string]Component)

	e := echo.New()

	e.Pre(rewriteFrontEndPaths())
	e.Use(middleware.CORS())

	e.GET("/api", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/api/inbox", withCache(getInbox, cache))

	e.Static("/", "/app/web")

	go populateComponentCache(cache)

	e.Logger.Fatal(e.Start(":1323"))
}

func getInbox(c echo.Context, cache ComponentCache) error {
	components := make([]ComponentStatus, 0, len(cache))
	for _, cacheItem := range cache {
		status := "unknown"
		if secondsSinceUpdate(cacheItem) < 10 {
			status = "online"
		}

		component := ComponentStatus{
			Uuid:   cacheItem.Uuid,
			Status: status,
		}

		components = append(components, component)
	}

	return c.JSON(http.StatusOK, components)
}
