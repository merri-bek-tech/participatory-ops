package main

import (
	"context"
	comms "paropd/comms"
	config "paropd/config/computed"

	"github.com/google/uuid"
)

func main() {
	sampleConfig, err := config.LoadFromPath(context.Background(), "./config/computed/sample.pkl")
	if err != nil {
		panic(err)
	} else {
		println("Hostname from config: ", sampleConfig.HostName)
	}

	deviceId := uuid.New().String()

	client := comms.Connect(deviceId)

	client.PublishHeartbeat()

	client.Disconnect()
}
