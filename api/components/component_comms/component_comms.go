package component_comms

import (
	"parops.libs/msg"

	uuid "github.com/google/uuid"
)

func MonitorComponents(handlers msg.CommsHandlers) {
	deviceId := "api-" + uuid.New().String()
	client := msg.Connect(deviceId)

	client.Subscribe("components/+", handlers)
}
