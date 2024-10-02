package main

import (
	"log"
	"paropd/client"
	configs "paropd/config"
	"paropd/telemetry"

	"parops.libs/msg"
)

type ConnectedApp struct {
	config *configs.Config
	client *client.PahoConnection
}

func StartConnectedApp(config *configs.Config) *ConnectedApp {
	result := &ConnectedApp{config: config}
	params := client.MqttConnectionParams{
		Host: "127.0.0.1",
		Port: 1883,
	}

	result.client = client.Connect(config.Computed.Uuid, params)

	log.Println("Broker found, starting connection")

	result.startSubscriptions()

	return result
}

func (app *ConnectedApp) Close() {
	if app.client != nil {
		log.Println("Closing client connection")
		app.client.Disconnect()
		app.client = nil
	}
}

func (app *ConnectedApp) HeartbeatTick() {
	app.client.GetMessenger().PublishMyHeartbeat(app.config.SchemeId)
}

func (app *ConnectedApp) startSubscriptions() {
	handlers := msg.CommsHandlers{
		HandleHeartbeat:  nil,
		DetailsRequested: func(schemeId string) { app.onDetailsRequested(schemeId) },
	}
	app.client.GetMessenger().SubscribeDevice(app.config.SchemeId, handlers)

}

func (app *ConnectedApp) onDetailsRequested(schemeId string) {
	log.Println("Details requested")

	if app.config.SchemeId != schemeId {
		log.Printf("[%s] Scheme ID mismatch. Got %s\n", app.config.SchemeId, schemeId)
		return
	}

	telemetry.SendDetails(app.config, app.client)
}
