package main

import (
	"log"
	"net/http"
	"parops/broker"
	comps "parops/components"
	compCache "parops/components/component_cache"
	"parops/schemes"

	msg "parops.libs/msg"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	caches := compCache.NewComponentCachesForSchemes()

	e := echo.New()

	e.Pre(rewriteFrontEndPaths())
	e.Use(middleware.CORS())

	e.GET("/api", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/api/schemes", schemes.GetIndex)
	e.GET("/api/schemes/:schemeId/inbox", schemes.WithScheme(compCache.WithCache(comps.GetInbox, caches)))

	e.Static("/", "/app/web")

	go broker.MessageBroker(func(messenger *msg.Messenger) {
		log.Println("Broker onStarted")
		go comps.MonitorComponents(caches, messenger)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
