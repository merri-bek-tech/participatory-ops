package components

import (
	"fmt"
	"parops.libs/messages"
	compCache "parops/components/component_cache"
	comms "parops/components/component_comms"
)

var detailsCheckFrequencySeconds int64 = 20

func MonitorComponents(cache *compCache.ComponentCache) {
	comms.MonitorComponents(comms.CommsHandlers{
		HandleHeartbeat: func(heartbeat messages.ComponentHeartbeat) {
			OnHeartbeat(heartbeat, cache)
		}})
}

func OnHeartbeat(heartbeat messages.ComponentHeartbeat, cache *compCache.ComponentCache) {
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
