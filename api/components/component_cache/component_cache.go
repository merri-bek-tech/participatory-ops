package component_cache

import (
	"fmt"
	events "parops/components/component_events"
	"time"

	"github.com/labstack/echo/v4"
)

type ComponentDetails struct {
	HostName string `json:"hostName"`
}

type Component struct {
	Uuid               string            `json:"uuid"`
	UpdatedAt          int64             `json:"at"`
	Details            *ComponentDetails `json:"details"`
	DetailsRequestedAt int64
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
	if secondsSince(component.UpdatedAt) < 10 {
		status = "online"
	}
	return status
}

func OnHeartbeat(heartbeat events.ComponentHeartbeat, cache ComponentCache) {
	existing, exists := cache[heartbeat.Uuid]

	if !exists {
		cache[heartbeat.Uuid] = Component{
			Uuid:      heartbeat.Uuid,
			UpdatedAt: heartbeat.At,
		}
	} else {
		existing.UpdatedAt = heartbeat.At
	}
}

func (cache ComponentCache) FetchComponent(uuid string) (*Component, bool) {
	component, exists := cache[uuid]
	return &component, exists
}

func (component *Component) NeedsDetails(minCheckSeconds int64) bool {
	fmt.Println("Details last requested at: ", component.DetailsRequestedAt)

	return component.Details == nil && secondsSince(component.DetailsRequestedAt) > minCheckSeconds
}

func (component *Component) DetailsRequested() {
	component.DetailsRequestedAt = time.Now().Unix()
	fmt.Printf("detailsRequestedAt updated: %v\n", component)
}

// PRIVATE

func secondsSince(timestamp int64) int64 {
	return time.Now().Unix() - timestamp
}
