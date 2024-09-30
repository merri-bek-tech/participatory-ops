package broker

import (
	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/packets"
	msg "parops.libs/msg"
)

type publishFunc func(topic string, payload string)
type subscribeFunc func(topicFilter string, handler msg.SubscribeHandler)

type InlineClient struct {
	Publish   publishFunc
	Subscribe subscribeFunc
}

func BuildInlineClient(server *mqtt.Server) *InlineClient {
	return &InlineClient{
		Publish: func(topic string, payload string) {
			server.Publish(topic, []byte(payload), false, 0)
		},
		Subscribe: func(topicFilter string, handler msg.SubscribeHandler) {
			callbackFn := func(cl *mqtt.Client, sub packets.Subscription, pk packets.Packet) {
				server.Log.Info("inline client received message from subscription", "client", cl.ID, "subscriptionId", sub.Identifier, "topic", pk.TopicName, "payload", string(pk.Payload))
				handler(pk.TopicName, string(pk.Payload))
			}
			server.Subscribe(topicFilter, 1, callbackFn)
		},
	}
}
