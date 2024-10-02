package main

import (
	"log"

	"paropd/client"
	configs "paropd/config"
	"paropd/telemetry"
	"parops.libs/msg"
)

type AppData struct {
	config *configs.Config
	client *client.PahoConnection
}

func (app *AppData) init() error {
	app.close()

	log.Println("Initializing app")
	app.config = configs.LoadConfig(true)
	app.client = client.Connect(app.config.Computed.Uuid)

	handlers := msg.CommsHandlers{
		HandleHeartbeat:  nil,
		DetailsRequested: func(schemeId string) { app.onDetailsRequested(schemeId) },
	}
	app.client.GetMessenger().SubscribeDevice(app.config.SchemeId, handlers)

	return nil
}

func (app *AppData) close() {
	if app.client != nil {
		log.Println("Closing client connection")
		app.client.Disconnect()
		app.client = nil
	}

	if app.config != nil {
		app.config = nil
	}
}

func (app *AppData) onDetailsRequested(schemeId string) {
	log.Println("Details requested")

	if app.config.SchemeId != schemeId {
		log.Printf("[%s] Scheme ID mismatch. Got %s\n", app.config.SchemeId, schemeId)
		return
	}

	telemetry.SendDetails(app.config, app.client)
}
