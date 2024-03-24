package main

import (
	"fmt"
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

	go componentCache(cache)

	e.Logger.Fatal(e.Start(":1323"))
}

func withCache(next func(c echo.Context, cache map[string]Component) error, cache map[string]Component) echo.HandlerFunc {
	return func(context echo.Context) error {
		return next(context, cache)
	}
}

func getInbox(c echo.Context, cache map[string]Component) error {
	components := make([]ComponentStatus, 0, len(cache))
	for _, cacheItem := range cache {
		fmt.Println("mapping cacheItem", cacheItem)

		component := ComponentStatus{
			Uuid:   cacheItem.Uuid,
			Status: "unknown",
		}

		components = append(components, component)
	}

	return c.JSON(http.StatusOK, components)
}
