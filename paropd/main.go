package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"syscall"
	"time"
)

const defaultTick = 6 * time.Second

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGHUP)

	log.SetOutput(os.Stdout)

	app := &AppData{}

	defer func() {
		signal.Stop(signalChan)
		app.Close()
		cancel()
	}()

	go func() {
		for {
			select {
			case s := <-signalChan:
				switch s {
				case syscall.SIGHUP:
					err := app.init()
					if err != nil {
						log.Println("Error initializing app", err)
					}
				case os.Interrupt:
					log.Println("Interrupted.")
					app.Close()
					cancel()
					os.Exit(1)
				}
			case <-ctx.Done():
				os.Exit(1)
			}
		}
	}()

	if err := run(ctx, app); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, app *AppData) error {
	err := app.init()
	if err != nil {
		log.Println("Error running app", err)
		return err
	}

	for {
		select {
		case <-ctx.Done():
			log.Println("Context done in the run function")
			return nil
		case <-time.Tick(defaultTick):
			app.HeartbeatTick()
		}
	}
}
