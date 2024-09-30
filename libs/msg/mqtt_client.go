package msg

type SubscribeHandler func(topic string, payload string)

type MqttClient interface {
	Publish(topic string, payload string)
	Subscribe(topicFilter string, handler SubscribeHandler)
}
