package component_cache

import (
	events "parops/component_events"
	"testing"
)

func TestCacheStartsEmpty(t *testing.T) {
	cache := NewComponentCache()
	if len(cache) != 0 {
		t.Error("Expected cache to start empty")
	}
}

func TestOnHeartbeatWithNewComponent(t *testing.T) {
	cache := NewComponentCache()
	heartbeat := events.ComponentHeartbeat{
		Uuid: "f08b7172-36d8-447f-85e1-41403d2730c8",
		At:   1234567890,
	}

	OnHeartbeat(heartbeat, cache)

	if len(cache) != 1 {
		t.Error("Expected cache to have one item")
	}
	if cache[heartbeat.Uuid].Uuid != heartbeat.Uuid {
		t.Error("Expected cache to have the correct component uuid")
	}
	if cache[heartbeat.Uuid].At != heartbeat.At {
		t.Error("Expected cache to have the correct timestamp")
	}
}

func TestOnHeartbeatWithSecondNewComponent(t *testing.T) {
	cache := NewComponentCache()
	heartbeat1 := events.ComponentHeartbeat{
		Uuid: "aaa",
		At:   1234567890,
	}
	heartbeat2 := events.ComponentHeartbeat{
		Uuid: "bbb",
		At:   1234567891,
	}

	OnHeartbeat(heartbeat1, cache)
	OnHeartbeat(heartbeat2, cache)

	if len(cache) != 2 {
		t.Error("Expected cache to have two items")
	}
}

func TestOnHeartbeatWithUpdateOnExistingComponent(t *testing.T) {
	cache := NewComponentCache()
	heartbeat1 := events.ComponentHeartbeat{
		Uuid: "aaa",
		At:   1234567890,
	}
	heartbeat2 := events.ComponentHeartbeat{
		Uuid: "aaa",
		At:   1234567891,
	}

	OnHeartbeat(heartbeat1, cache)
	OnHeartbeat(heartbeat2, cache)

	if len(cache) != 1 {
		t.Error("Expected cache to have one item")
	}
	if cache[heartbeat1.Uuid].At != heartbeat2.At {
		t.Error("Expected cache to have the updated timestamp")
	}
}

func TestOnHeartbeatWithOlderMessageAboutExistingComponent(t *testing.T) {
	cache := NewComponentCache()
	heartbeat1 := events.ComponentHeartbeat{
		Uuid: "aaa",
		At:   1234567891,
	}
	heartbeat2 := events.ComponentHeartbeat{
		Uuid: "aaa",
		At:   1234567890,
	}

	OnHeartbeat(heartbeat1, cache)
	OnHeartbeat(heartbeat2, cache)

	if len(cache) != 1 {
		t.Error("Expected cache to have one item")
	}
	if cache[heartbeat1.Uuid].At != heartbeat1.At {
		t.Error("Expected cache to have the original timestamp")
	}
}
