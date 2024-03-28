package component_cache

import (
	events "parops/component_events"
	"time"

	"github.com/labstack/echo/v4"
)

type Component struct {
	Uuid string `json:"uuid"`
	At   int64  `json:"at"`
}

type ComponentCache map[string]Component

type HandlerWithComponentCache func(c echo.Context, cache ComponentCache) error

func NewComponentCache() ComponentCache {
	return make(map[string]Component)
}

func WithCache(next HandlerWithComponentCache, cache ComponentCache) echo.HandlerFunc {
	return func(context echo.Context) error {
		return next(context, cache)
	}
}

func Status(component Component) string {
	status := "unknown"
	if secondsSinceUpdate(component) < 10 {
		status = "online"
	}
	return status
}

func OnHeartbeat(heartbeat events.ComponentHeartbeat, cache ComponentCache) {
	existing, exists := cache[heartbeat.Uuid]
	if !exists || existing.At < heartbeat.At {
		cache[heartbeat.Uuid] = Component{
			Uuid: heartbeat.Uuid,
			At:   heartbeat.At,
		}
	}
}

func PopulateComponentCache(cache ComponentCache) {
	time.Sleep(3 * time.Second)

	cache["f08b7172-36d8-447f-85e1-41403d2730c8"] = Component{
		Uuid: "f08b7172-36d8-447f-85e1-41403d2730c8",
		At:   time.Now().Unix(),
	}
}

// PRIVATE

func secondsSinceUpdate(component Component) int64 {
	return time.Now().Unix() - component.At
}
