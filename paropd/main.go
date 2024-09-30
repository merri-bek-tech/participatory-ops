package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"syscall"
	"time"

	configs "paropd/config"

	"parops.libs/msg"
)

const defaultTick = 6 * time.Second

type AppData struct {
	config *configs.Config
	client *msg.PahoConnection
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGHUP)

	log.SetOutput(os.Stdout)

	app := &AppData{}

	defer func() {
		signal.Stop(signalChan)
		app.close()
		cancel()
	}()

	go func() {
		for {
			select {
			case s := <-signalChan:
				switch s {
				case syscall.SIGHUP:
					app.init()
				case os.Interrupt:
					log.Println("Interrupted.")
					app.close()
					cancel()
					os.Exit(1)
				}
			case <-ctx.Done():
				log.Printf("Done.")
				os.Exit(1)
			}
		}
	}()

	if err := run(ctx, app); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func (app *AppData) init() error {
	app.close()

	log.Println("Initializing app")
	app.config = configs.LoadConfig(true)
	app.client = msg.Connect(app.config.Computed.Uuid)

	handlers := msg.CommsHandlers{
		HandleHeartbeat:  nil,
		DetailsRequested: func(schemeId string) { app.onDetailsRequested(schemeId) },
	}
	app.client.SubscribeDevice(app.config.SchemeId, handlers)

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

func run(ctx context.Context, app *AppData) error {
	app.init()

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-time.Tick(defaultTick):
			app.client.GetMessenger().PublishMyHeartbeat(app.config.SchemeId)
		}
	}
}

func (app *AppData) onDetailsRequested(schemeId string) {
	log.Println("Details requested")

	if app.config.SchemeId != schemeId {
		log.Printf("[%s] Scheme ID mismatch. Got %s\n", app.config.SchemeId, schemeId)
		return
	}

	details := msg.ComponentDetails{
		Uuid:        app.config.Computed.Uuid,
		HostName:    app.config.Computed.HostName,
		ProductName: app.config.Computed.ProductName,
		SysVendor:   app.config.Computed.SysVendor,
	}

	log.Printf("Publishing details: %v\n", details)

	app.client.PublishDetails(app.config.SchemeId, app.config.Computed.Uuid, details)
}
