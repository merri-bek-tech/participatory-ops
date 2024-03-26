package main

import (
	"fmt"
	"net/http"
	cache "parops/component_cache"
	comms "parops/component_comms"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ComponentStatus struct {
	Uuid   string `json:"uuid"`
	Status string `json:"status"`
}

func main() {
	cacheData := make(map[string]cache.Component)

	e := echo.New()

	e.Pre(rewriteFrontEndPaths())
	e.Use(middleware.CORS())

	e.GET("/api", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/api/inbox", cache.WithCache(getInbox, cacheData))

	e.Static("/", "/app/web")

	go cache.PopulateComponentCache(cacheData)
	go comms.MonitorComponents(comms.CommsHandlers{
		HandleHeartbeat: handleHeartbeat,
	})

	e.Logger.Fatal(e.Start(":1323"))
}

func handleHeartbeat(heartbeat comms.ComponentHeartbeat) {
	fmt.Println("Received heartbeat: ", heartbeat)
}

func getInbox(c echo.Context, cacheData cache.ComponentCache) error {
	components := make([]ComponentStatus, 0, len(cacheData))
	for _, cacheItem := range cacheData {
		component := ComponentStatus{
			Uuid:   cacheItem.Uuid,
			Status: cache.Status(cacheItem),
		}

		components = append(components, component)
	}

	return c.JSON(http.StatusOK, components)
}
