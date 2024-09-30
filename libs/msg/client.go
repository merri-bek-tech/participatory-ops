package msg

type SubscribeHandler func(topic string, payload string)

type Client interface {
	Publish(topic string, payload string)
	Subscribe(topicFilter string, handler SubscribeHandler)
}
