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

const defaultTick = 3 * time.Second

type AppData struct {
	config *configs.Config
	client *msg.Client
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
	app.client.SubscribeDevice(msg.CommsHandlers{
		HandleHeartbeat:  nil,
		DetailsRequested: func(schemeId string) { app.onDetailsRequested(schemeId) },
	})

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
			app.client.PublishMyHeartbeat(app.config.SchemeId)
		}
	}
}

func (app *AppData) onDetailsRequested(schemeId string) {
	log.Println("Details requested")

	if app.config.SchemeId != schemeId {
		log.Printf("[%s] Scheme ID mismatch. Got %s\n", app.config.SchemeId, schemeId)
		return
	}

	app.client.PublishDetails(app.config.Computed.Uuid, msg.ComponentDetails{
		Uuid:     app.config.Computed.Uuid,
		HostName: app.config.Computed.HostName,
	})
}
