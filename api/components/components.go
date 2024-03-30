package components

import (
	compCache "parops/components/component_cache"
	comms "parops/components/component_comms"
	events "parops/components/component_events"
)

var detailsCheckFrequencySeconds int64 = 20

func MonitorComponents(cache *compCache.ComponentCache) {
	comms.MonitorComponents(comms.CommsHandlers{
		HandleHeartbeat: func(heartbeat events.ComponentHeartbeat) {
			cache.OnHeartbeat(heartbeat)
		}})
}

func OnHeartbeat(heartbeat events.ComponentHeartbeat, cache *compCache.ComponentCache) {
	// component1, exists1 := cache.Get(heartbeat.Uuid)
	// if exists1 {
	// 	fmt.Printf(" - Component exists 1 %v\n", component1)
	// }

	cache.OnHeartbeat(heartbeat)

	// component, exists := cache.Get(heartbeat.Uuid)
	// if exists {
	// 	fmt.Printf(" - Component exists 2 %v\n", component)
	// 	if component.NeedsDetails(detailsCheckFrequencySeconds) {
	// 		RequestDetails(component)
	// 	}
	// }

	// component3, exists3 := cacheData.FetchComponent(heartbeat.Uuid)
	// if exists3 {
	// 	fmt.Printf(" - Component exists 3 %v\n", component3)
	// }
}

// func RequestDetails(component *compCache.Component) {
// 	fmt.Printf("Requesting details for %s\n", component.Uuid)

// 	component.DetailsRequested()
// }
