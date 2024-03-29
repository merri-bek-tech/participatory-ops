package main

import (
	comms "paropd/comms"

	"github.com/google/uuid"
)

func main() {
	deviceId := uuid.New().String()

	client := comms.Connect(deviceId)

	client.PublishHeartbeat()

	client.Disconnect()
}
