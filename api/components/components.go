package components

import (
	"fmt"
	compCache "parops/components/component_cache"
	comms "parops/components/component_comms"

	"parops.libs/msg"
)

var detailsCheckFrequencySeconds int64 = 20

func MonitorComponents(cache *compCache.ComponentCache) {
	comms.MonitorComponents(msg.CommsHandlers{
		HandleHeartbeat: func(heartbeat msg.ComponentHeartbeat) {
			OnHeartbeat(heartbeat, cache)
		}})
}

func OnHeartbeat(heartbeat msg.ComponentHeartbeat, cache *compCache.ComponentCache) {
	cache.OnHeartbeat(heartbeat)

	component, exists := cache.Get(heartbeat.Uuid)
	if exists {
		if component.NeedsDetails(detailsCheckFrequencySeconds) {
			RequestDetails(component)
		}
	}
}

func RequestDetails(component *compCache.Component) {
	fmt.Printf("Requesting details for %s\n", component.Uuid)

	component.DetailsRequested()
}
