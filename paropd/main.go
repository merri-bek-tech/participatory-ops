package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"

	"syscall"
	"time"

	comms "paropd/comms"
	configs "paropd/config"
)

const defaultTick = 3 * time.Second

// func main() {
// 	// configData := config.LoadConfig(true)
// 	// fmt.Printf("Loaded config: %+v\n", configData.Computed)

// 	// client := comms.Connect(configData.Computed.Uuid)

// 	// client.PublishHeartbeat()

// 	// client.Disconnect()

// 	ctx := context.Background()
// 	ctx, cancel := context.WithCancel(ctx)

// 	defer func() {
// 		cancel()
// 	}()

// 	if err := run(ctx); err != nil {
// 		fmt.Println("About to exit")
// 		fmt.Fprintf(os.Stderr, "%s\n", err)
// 		os.Exit(1)
// 	}
// }

// type config struct {
// 	contentType string
// 	server      string
// 	statusCode  int
// 	tick        time.Duration
// 	url         string
// 	userAgent   string
// }

type AppData struct {
	config *configs.Config
	client *comms.Client
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGHUP)

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
					fmt.Println("Interrupted.")
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

	if err := run(ctx, app, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func (app *AppData) init() error {
	app.close()

	log.Println("Initializing app")
	app.config = configs.LoadConfig(true)
	app.client = comms.Connect(app.config.Computed.Uuid)

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

func run(ctx context.Context, app *AppData, stdout io.Writer) error {
	app.init()
	log.SetOutput(stdout)

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-time.Tick(defaultTick):
			app.client.PublishHeartbeat()
		}
	}
}
