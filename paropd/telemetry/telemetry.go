package telemetry

import (
	"log"
	"paropd/client"
	configs "paropd/config"

	"parops.libs/msg"
)

func SendDetails(config *configs.Config, client *client.PahoConnection) {
	details := msg.ComponentDetails{
		Uuid:        config.Computed.Uuid,
		HostName:    config.Computed.HostName,
		ProductName: config.Computed.ProductName,
		SysVendor:   config.Computed.SysVendor,
	}

	log.Printf("Publishing details: %v\n", details)

	client.GetMessenger().PublishDetails(config.SchemeId, config.Computed.Uuid, details)
}
