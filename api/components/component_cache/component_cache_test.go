package component_cache

import (
	"testing"

	"parops.libs/paroplib/messages"
)

func TestCacheStartsEmpty(t *testing.T) {
	cache := NewComponentCache()
	if cache.ItemCount() != 0 {
		t.Error("Expected cache to start empty")
	}
}

func TestOnHeartbeatWithNewComponent(t *testing.T) {
	cache := NewComponentCache()
	heartbeat := messages.ComponentHeartbeat{
		Uuid: "f08b7172-36d8-447f-85e1-41403d2730c8",
		At:   1234567890,
	}

	cache.OnHeartbeat(heartbeat)

	if cache.ItemCount() != 1 {
		t.Error("Expected cache to have one item")
	}

	existing, exists := cache.Get(heartbeat.Uuid)

	if !exists {
		t.Error("Expected cache to have the new component")
	} else {
		if existing.Uuid != heartbeat.Uuid {
			t.Error("Expected cache to have the correct component uuid")
		}
		if existing.UpdatedAt != heartbeat.At {
			t.Error("Expected cache to have the correct timestamp")
		}
	}
}

func TestOnHeartbeatWithSecondNewComponent(t *testing.T) {
	cache := NewComponentCache()
	heartbeat1 := messages.ComponentHeartbeat{
		Uuid: "aaa",
		At:   1234567890,
	}
	heartbeat2 := messages.ComponentHeartbeat{
		Uuid: "bbb",
		At:   1234567891,
	}

	cache.OnHeartbeat(heartbeat1)
	cache.OnHeartbeat(heartbeat2)

	if cache.ItemCount() != 2 {
		t.Error("Expected cache to have two items")
	}
}

func TestOnHeartbeatWithUpdateOnExistingComponent(t *testing.T) {
	cache := NewComponentCache()
	heartbeat1 := messages.ComponentHeartbeat{
		Uuid: "aaa",
		At:   1234567890,
	}
	heartbeat2 := messages.ComponentHeartbeat{
		Uuid: "aaa",
		At:   1234567891,
	}

	cache.OnHeartbeat(heartbeat1)
	cache.OnHeartbeat(heartbeat2)

	if cache.ItemCount() != 1 {
		t.Error("Expected cache to have one item")
	}

	existing, exists := cache.Get("aaa")
	if !exists {
		t.Error("Expected cache to have the new component")
	} else {
		if existing.Uuid != heartbeat2.Uuid {
			t.Error("Expected cache to have the updated timestamp")
		}
	}
}

func TestOnHeartbeatWithOlderMessageAboutExistingComponent(t *testing.T) {
	cache := NewComponentCache()
	heartbeat1 := messages.ComponentHeartbeat{
		Uuid: "aaa",
		At:   1234567891,
	}
	heartbeat2 := messages.ComponentHeartbeat{
		Uuid: "aaa",
		At:   1234567890,
	}

	cache.OnHeartbeat(heartbeat1)
	cache.OnHeartbeat(heartbeat2)

	if cache.ItemCount() != 1 {
		t.Error("Expected cache to have one item")
	}

	existing, exists := cache.Get("aaa")
	if !exists {
		t.Error("Expected cache to have the new component")
	} else {
		if existing.Uuid != heartbeat1.Uuid {
			t.Error("Expected cache to have the original timestamp")
		}
	}
}

func TestItemListReturnsAnArrayOfItems(t *testing.T) {
	cache := NewComponentCache()
	cache.OnHeartbeat(messages.ComponentHeartbeat{Uuid: "aaa", At: 1234567891})
	cache.OnHeartbeat(messages.ComponentHeartbeat{Uuid: "bbb", At: 1234567891})

	items := cache.ItemList()

	if len(items) != 2 {
		t.Error("Expected cache to have two items")
	}
	if items[0].Uuid != "aaa" {
		t.Error("Expected first item to have uuid 'aaa' but was ", items[0].Uuid)
	}
	if items[1].Uuid != "bbb" {
		t.Error("Expected second item to have uuid 'bbb' but was ", items[1].Uuid)
	}
}
