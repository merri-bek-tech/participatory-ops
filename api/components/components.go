package components

import (
	"fmt"
	"log"
	compCache "parops/components/component_cache"

	"github.com/google/uuid"
	"parops.libs/msg"
)

var detailsCheckFrequencySeconds int64 = 20

func MonitorComponents(cache *compCache.ComponentCache) {
	deviceId := "api-" + uuid.New().String()
	client := msg.Connect(deviceId)
	handlers := msg.CommsHandlers{
		HandleHeartbeat: func(heartbeat msg.ComponentHeartbeat) {
			OnHeartbeat(heartbeat, cache, client)
		},
		ComponentDetails: func(details msg.ComponentDetails) {
			log.Printf("Received details for %s: %s\n", details.Uuid, details.HostName)
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
	fmt.Printf("Requesting details for %s\n", component.Uuid)

	client.PublishDetailsRequested(component.Uuid)
	component.DetailsRequested()
}
