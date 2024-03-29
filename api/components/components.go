package components

import (
	cache "parops/components/component_cache"
	comms "parops/components/component_comms"
	events "parops/components/component_events"
)

func MonitorComponents(cacheData cache.ComponentCache) {
	comms.MonitorComponents(comms.CommsHandlers{
		HandleHeartbeat: func(heartbeat events.ComponentHeartbeat) {
			OnHeartbeat(heartbeat, cacheData)
		}})
}

func OnHeartbeat(heartbeat events.ComponentHeartbeat, cacheData cache.ComponentCache) {
	cache.OnHeartbeat(heartbeat, cacheData)
}
