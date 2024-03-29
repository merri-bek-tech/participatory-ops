package main

import (
	"fmt"
	"paropd/comms"
	config "paropd/config"
)

func main() {
	configData := config.LoadConfig(true)
	fmt.Printf("Loaded config: %+v\n", configData.Computed)

	client := comms.Connect(configData.Computed.Uuid)

	client.PublishHeartbeat()

	client.Disconnect()
}
