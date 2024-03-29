package main

import (
	"fmt"
	comms "paropd/comms"
	config "paropd/config"

	"github.com/google/uuid"
)

func main() {
	configData, err := config.LoadConfig(false)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("Loaded config: %+v\n", configData.Computed)
	}

	deviceId := uuid.New().String()

	client := comms.Connect(deviceId)

	client.PublishHeartbeat()

	client.Disconnect()
}
