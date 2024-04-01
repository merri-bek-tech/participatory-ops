package components

import (
	"log"
	"net/http"
	compCache "parops/components/component_cache"
	"parops/schemes"

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

func GetInbox(c echo.Context, _ *schemes.SchemeIdentity, cache *compCache.ComponentCache) error {
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

func MonitorComponents(caches *map[string]*compCache.ComponentCache) {
	deviceId := "api-" + uuid.New().String()
	client := msg.Connect(deviceId)

	handlers := msg.CommsHandlers{
		HandleHeartbeat: func(schemeId string, heartbeat msg.ComponentHeartbeat) {
			log.Printf("[%s] received heartbeat\n", schemeId)
			cache := compCache.CacheForScheme(caches, schemeId)
			OnHeartbeat(schemeId, heartbeat, cache, client)
		},
		ComponentDetails: func(schemeId string, details msg.ComponentDetails) {
			log.Printf("Received details (%s) for %s: %s\n", schemeId, details.Uuid, details.HostName)
			cache := compCache.CacheForScheme(caches, schemeId)
			cache.SetDetails(details.Uuid, &details)
		},
	}

	client.SubscribeAllComponents(handlers)
}

func OnHeartbeat(schemeId string, heartbeat msg.ComponentHeartbeat, cache *compCache.ComponentCache, client *msg.Client) {
	cache.OnHeartbeat(heartbeat)

	component, exists := cache.Get(heartbeat.Uuid)
	if exists {
		if component.NeedsDetails(detailsCheckFrequencySeconds) {
			RequestDetails(schemeId, component, client)
		}
	}
}

func RequestDetails(schemeId string, component *compCache.Component, client *msg.Client) {
	log.Printf("Requesting details for %s\n", component.Uuid)

	client.PublishDetailsRequested(schemeId, component.Uuid)
	component.DetailsRequested()
}
