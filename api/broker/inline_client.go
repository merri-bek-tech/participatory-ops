package broker

import (
	mqtt "github.com/mochi-mqtt/server/v2"
)

type publishFunc func(topic string, payload []byte, retain bool, qos byte) error

type InlineClient struct {
	Publish publishFunc
}

func BuildInlineClient(server *mqtt.Server) *InlineClient {
	return &InlineClient{
		Publish: func(topic string, payload []byte, retain bool, qos byte) error {
			return server.Publish(topic, payload, retain, qos)
		},
	}
}
