package components

import (
	"log"
	"net/http"
	compCache "parops/components/component_cache"

	"github.com/labstack/echo/v4"

	"github.com/google/uuid"
	"parops.libs/msg"
)

var detailsCheckFrequencySeconds int64 = 20

type ComponentStatus struct {
	Uuid    string                `json:"uuid"`
	Status  string                `json:"status"`
	Details *msg.ComponentDetails `json:"details"`
}

func GetInbox(c echo.Context, cache *compCache.ComponentCache) error {
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

func MonitorComponents(cache *compCache.ComponentCache) {
	deviceId := "api-" + uuid.New().String()
	client := msg.Connect(deviceId)
	handlers := msg.CommsHandlers{
		HandleHeartbeat: func(heartbeat msg.ComponentHeartbeat) {
			OnHeartbeat(heartbeat, cache, client)
		},
		ComponentDetails: func(details msg.ComponentDetails) {
			log.Printf("Received details for %s: %s\n", details.Uuid, details.HostName)
			cache.SetDetails(details.Uuid, &details)
		},
	}

	client.SubscribeAllComponents(handlers)
}

func OnHeartbeat(heartbeat msg.ComponentHeartbeat, cache *compCache.ComponentCache, client *msg.Client) {
	cache.OnHeartbeat(heartbeat)

	component, exists := cache.Get(heartbeat.Uuid)
	if exists {
		if component.NeedsDetails(detailsCheckFrequencySeconds) {
			RequestDetails(component, client)
		}
	}
}

func RequestDetails(component *compCache.Component, client *msg.Client) {
	log.Printf("Requesting details for %s\n", component.Uuid)

	client.PublishDetailsRequested(component.Uuid)
	component.DetailsRequested()
}
