package main

import (
	"log"
	configs "paropd/config"
)

type AppData struct {
	config       *configs.Config
	connectedApp *ConnectedApp
}

func (app *AppData) init() error {
	app.Close()

	log.Println("Initializing app")
	app.config = configs.LoadConfig(true)

	app.connectedApp = StartConnectedApp(app.config)

	return nil
}

func (app *AppData) Close() {
	if app.connectedApp != nil {
		app.connectedApp.Close()
		app.connectedApp = nil
	}

	if app.config != nil {
		app.config = nil
	}

	if app.connectedApp != nil {
		app.connectedApp = nil
	}
}

func (app *AppData) HeartbeatTick() {
	if app.connectedApp != nil {
		app.connectedApp.HeartbeatTick()
	}
}
