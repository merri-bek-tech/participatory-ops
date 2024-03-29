package main

import (
	config "paropd/config"
)

func main() {
	config.LoadConfig(false)
	// configData, err := config.LoadConfig(false)
	// if err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Printf("Loaded config: %+v\n", configData.Computed)
	// }

	// deviceId := uuid.New().String()

	// client := comms.Connect(deviceId)

	// client.PublishHeartbeat()

	// client.Disconnect()
}
