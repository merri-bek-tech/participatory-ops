package main

import (
	"context"
	"fmt"
	comms "paropd/comms"
	config "paropd/config/computed"

	"github.com/google/uuid"
)

func main() {
	sampleConfig, err := config.LoadFromPath(context.Background(), "./config/computed/sample.pkl")
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("Loaded config: %+v\n", sampleConfig)
	}

	deviceId := uuid.New().String()

	client := comms.Connect(deviceId)

	client.PublishHeartbeat()

	client.Disconnect()
}
