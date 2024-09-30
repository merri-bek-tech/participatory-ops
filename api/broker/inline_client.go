package broker

import (
	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/packets"
)

type genericMessageHandler func(topic string, payload []byte)
type publishFunc func(topic string, payload []byte, retain bool, qos byte) error
type subscribeFunc func(topic string, qos byte, handler genericMessageHandler) error

type InlineClient struct {
	Publish   publishFunc
	Subscribe subscribeFunc
}

func BuildInlineClient(server *mqtt.Server) *InlineClient {
	return &InlineClient{
		Publish: func(topic string, payload []byte, retain bool, qos byte) error {
			return server.Publish(topic, payload, retain, qos)
		},
		Subscribe: func(topic string, qos byte, handler genericMessageHandler) error {
			callbackFn := func(cl *mqtt.Client, sub packets.Subscription, pk packets.Packet) {
				server.Log.Info("inline client received message from subscription", "client", cl.ID, "subscriptionId", sub.Identifier, "topic", pk.TopicName, "payload", string(pk.Payload))
				handler(pk.TopicName, pk.Payload)
			}
			return server.Subscribe(topic, 0, callbackFn)
		},
	}
}
