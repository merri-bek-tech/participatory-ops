package component_cache

import (
	"cmp"
	"log"
	"parops/schemes"
	"slices"
	"time"

	"parops.libs/msg"

	"github.com/labstack/echo/v4"
	gocaches "github.com/patrickmn/go-cache"
)

type ComponentCache struct {
	gocache *gocaches.Cache
}

type Component struct {
	Uuid               string                `json:"uuid"`
	UpdatedAt          int64                 `json:"at"`
	Details            *msg.ComponentDetails `json:"details"`
	DetailsRequestedAt int64
}

type HandlerWithComponentCache func(c echo.Context, scheme *schemes.SchemeIdentity, cache *ComponentCache) error

func NewComponentCache() *ComponentCache {
	return &ComponentCache{
		gocache: gocaches.New(gocaches.NoExpiration, 10*time.Minute),
	}
}

func NewComponentCachesForSchemes() *map[string]*ComponentCache {
	caches := map[string]*ComponentCache{}
	return &caches
}

func WithCache(next HandlerWithComponentCache, caches *map[string]*ComponentCache) schemes.HandlerFuncWithScheme {
	return func(context echo.Context, scheme *schemes.SchemeIdentity) error {
		return next(context, scheme, CacheForScheme(caches, scheme.Id))
	}
}

func CacheForScheme(caches *map[string]*ComponentCache, schemeId string) *ComponentCache {
	if schemeId == "" {
		log.Println("schemeId is required, using transient cache")
		NewComponentCache()
	}

	if cache, exists := (*caches)[schemeId]; !exists {
		newCache := NewComponentCache()
		(*caches)[schemeId] = newCache
		return newCache
	} else {
		return cache
	}
}

func Status(component *Component) string {
	//log.Printf("Checking status for %v\n", component)

	status := "unknown"
	if secondsSince(component.UpdatedAt) < 10 {
		status = "online"
	}
	return status
}

func (cache *ComponentCache) OnHeartbeat(heartbeat msg.ComponentHeartbeat) {
	existing, exists := cache.gocache.Get(heartbeat.Uuid)

	if !exists {
		newComponent := Component{
			Uuid:      heartbeat.Uuid,
			UpdatedAt: heartbeat.At,
		}
		cache.gocache.SetDefault(heartbeat.Uuid, &newComponent)
	} else {
		existing.(*Component).UpdatedAt = heartbeat.At
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

func (cache *ComponentCache) SetDetails(uuid string, details *msg.ComponentDetails) {
	component, exists := cache.Get(uuid)
	if exists {
		component.Details = details
	} else {
		log.Printf("Attempted to set details for unknown component: %s\n", uuid)
	}
}

func (component *Component) NeedsDetails(minCheckSeconds int64) bool {
	// log.Println("Details last requested at: ", component.DetailsRequestedAt)
	return component.Details == nil && secondsSince(component.DetailsRequestedAt) > minCheckSeconds
}

func (component *Component) DetailsRequested() {
	component.DetailsRequestedAt = time.Now().Unix()
	log.Printf("detailsRequestedAt updated: %v\n", component)
}

// // PRIVATE

func secondsSince(timestamp int64) int64 {
	return time.Now().Unix() - timestamp
}
