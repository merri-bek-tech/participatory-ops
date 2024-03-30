package component_cache

import (
	"cmp"
	"fmt"
	events "parops/components/component_events"
	"slices"
	"time"

	"github.com/labstack/echo/v4"
	gocaches "github.com/patrickmn/go-cache"
)

type ComponentCache struct {
	gocache *gocaches.Cache
}

type ComponentDetails struct {
	HostName string `json:"hostName"`
}

type Component struct {
	Uuid               string            `json:"uuid"`
	UpdatedAt          int64             `json:"at"`
	Details            *ComponentDetails `json:"details"`
	DetailsRequestedAt int64
}

type HandlerWithComponentCache func(c echo.Context, cache *ComponentCache) error

func NewComponentCache() *ComponentCache {
	return &ComponentCache{
		gocache: gocaches.New(gocaches.NoExpiration, 10*time.Minute),
	}
}

func WithCache(next HandlerWithComponentCache, cache *ComponentCache) echo.HandlerFunc {
	return func(context echo.Context) error {
		return next(context, cache)
	}
}

func Status(component *Component) string {
	status := "unknown"
	if secondsSince(component.UpdatedAt) < 10 {
		status = "online"
	}
	return status
}

func (cache *ComponentCache) OnHeartbeat(heartbeat events.ComponentHeartbeat) {
	existing, exists := cache.gocache.Get(heartbeat.Uuid)

	if !exists {
		newComponent := Component{
			Uuid:      heartbeat.Uuid,
			UpdatedAt: heartbeat.At,
		}
		cache.gocache.SetDefault(heartbeat.Uuid, &newComponent)
	} else {
		existing.(*Component).Uuid = heartbeat.Uuid
	}
}

func (cache *ComponentCache) ItemCount() int {
	return cache.gocache.ItemCount()
}

func (cache *ComponentCache) ItemList() []*Component {
	components := make([]*Component, 0, cache.ItemCount())
	for _, item := range cache.gocache.Items() {
		components = append(components, item.Object.(*Component))
	}

	// sort components by uuid
	slices.SortFunc(components, func(a, b *Component) int {
		return cmp.Compare(a.Uuid, b.Uuid)
	})

	return components
}

func (cache *ComponentCache) Get(uuid string) (*Component, bool) {
	component, exists := cache.gocache.Get(uuid)
	if !exists {
		return nil, false
	} else {
		return component.(*Component), true
	}
}

func (component *Component) NeedsDetails(minCheckSeconds int64) bool {
	fmt.Println("Details last requested at: ", component.DetailsRequestedAt)

	return component.Details == nil && secondsSince(component.DetailsRequestedAt) > minCheckSeconds
}

func (component *Component) DetailsRequested() {
	component.DetailsRequestedAt = time.Now().Unix()
	fmt.Printf("detailsRequestedAt updated: %v\n", component)
}

// // PRIVATE

func secondsSince(timestamp int64) int64 {
	return time.Now().Unix() - timestamp
}
