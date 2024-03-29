package components

import (
	cache "parops/components/component_cache"
	comms "parops/components/component_comms"
	events "parops/components/component_events"
)

func MonitorComponents(cacheData cache.ComponentCache) {
	comms.MonitorComponents(comms.CommsHandlers{
		HandleHeartbeat: func(heartbeat events.ComponentHeartbeat) {
			cache.OnHeartbeat(heartbeat, cacheData)
		}})
}
