package broker

import (
	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/packets"
	msg "parops.libs/msg"
)

type publishFunc func(topic string, payload string)
type subscribeFunc func(topicFilter string, handler msg.SubscribeHandler)

type InlineClient struct {
	server *mqtt.Server
}

func (inlineClient *InlineClient) Publish(topic string, payload string) {
	inlineClient.server.Publish(topic, []byte(payload), false, 0)
}

func (inlineClient *InlineClient) Subscribe(topicFilter string, handler msg.SubscribeHandler) {
	callbackFn := func(cl *mqtt.Client, sub packets.Subscription, pk packets.Packet) {
		inlineClient.server.Log.Info("inline client received message from subscription", "client", cl.ID, "subscriptionId", sub.Identifier, "topic", pk.TopicName, "payload", string(pk.Payload))
		handler(pk.TopicName, string(pk.Payload))
	}
	inlineClient.server.Subscribe(topicFilter, 1, callbackFn)
}

func BuildInlineClient(server *mqtt.Server) msg.MqttClient {
	return &InlineClient{server: server}
}
