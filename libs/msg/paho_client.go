package msg

import paho "github.com/eclipse/paho.mqtt.golang"

var (
	AtMostOnce  byte = 0
	AtLeastOnce byte = 1
	ExactlyOnce byte = 2
)

type PahoClient struct {
	Mqtt *paho.Client
}

func (pahoClient *PahoClient) Publish(topic string, payload string) {
	mqtt := *pahoClient.Mqtt
	mqtt.Publish(topic, AtMostOnce, false, payload)
}

func (pahoClient *PahoClient) Subscribe(topicFilter string, handler SubscribeHandler) {
	mqtt := *pahoClient.Mqtt
	mqtt.Subscribe(topicFilter, AtMostOnce, func(client paho.Client, message paho.Message) {
		handler(message.Topic(), string(message.Payload()))
	})
}
