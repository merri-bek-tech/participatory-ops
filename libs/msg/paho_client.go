package msg

import (
	"log"

	paho "github.com/eclipse/paho.mqtt.golang"
)

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
	token := mqtt.Publish(topic, AtMostOnce, false, payload)
	handleResult(token)
}

func (pahoClient *PahoClient) Subscribe(topicFilter string, handler SubscribeHandler) {
	mqtt := *pahoClient.Mqtt
	mqtt.Subscribe(topicFilter, AtMostOnce, func(client paho.Client, message paho.Message) {
		handler(message.Topic(), string(message.Payload()))
	})
}

// PRIVATE

func handleResult(token paho.Token) {
	token.Wait()
	// Check for errors during publishing (More on error reporting https://pkg.go.dev/github.com/eclipse/paho.mqtt.golang#readme-error-handling)
	if token.Error() != nil {
		log.Println("Failed to publish to topic")
		log.Println(token.Error())
	}
}
