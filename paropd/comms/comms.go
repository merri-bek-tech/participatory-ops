package comms

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	paho "github.com/eclipse/paho.mqtt.golang"
)

type Client struct {
	DeviceId string
	mqtt     paho.Client
}

type ComponentHeartbeat struct {
	Uuid string `json:"uuid"`
	At   int64  `json:"at"`
}

type Meta struct {
	Type    string `json:"type"`
	Version string `json:"version"`
}

func Connect(deviceId string) *Client {
	client := connectClient(
		"82e12caef57c4c8288d08fe23854c097.s1.eu.hivemq.cloud",
		8883,
		"paropd",
		"be9eiQuo",
		deviceId,
	)

	return &Client{
		DeviceId: deviceId,
		mqtt:     client,
	}
}

func (client *Client) Disconnect() {
	client.mqtt.Disconnect(250)
}

func (client *Client) PublishHeartbeat() {
	transmitHeartbeat("components/"+client.DeviceId, client.mqtt, client.DeviceId)
}

// PRIVATE

func transmitHeartbeat(topic string, client mqtt.Client, clientId string) mqtt.Token {
	log.Println("Publishing heartbeat")

	meta := Meta{
		Type:    "ComponentHeartbeat",
		Version: "1.0",
	}

	payload := ComponentHeartbeat{
		Uuid: clientId,
		At:   time.Now().Unix(),
	}

	metaText := jsonString(meta)
	payloadText := jsonString(payload)

	qos := byte(0) // 0 = at most once, 1 = at least once, 2 = exactly once

	return client.Publish(topic, qos, false, strings.Join([]string{metaText, payloadText}, "|"))
}

func connectClient(host string, port int, username string, password string, clientId string) paho.Client {
	opts := paho.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tls://%s:%d", host, port))
	opts.SetClientID(clientId) // set a name as you desire
	opts.SetUsername(username) // these are the credentials that you declare for your cluster (see readme)
	opts.SetPassword(password)
	// (optionally) configure callback handlers that get called on certain events
	// opts.SetDefaultPublishHandler(messagePubHandler)
	// opts.OnConnect = connectHandler
	// opts.OnConnectionLost = connectLostHandler
	// create the client using the options above
	client := paho.NewClient(opts)
	// throw an error if the connection isn't successfull
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	return client
}

func jsonString(input any) string {
	output, err := json.Marshal(input)
	if err != nil {
		panic(err)
	}
	return string(output)
}

func handleResult(token mqtt.Token) {
	token.Wait()
	// Check for errors during publishing (More on error reporting https://pkg.go.dev/github.com/eclipse/paho.mqtt.golang#readme-error-handling)
	if token.Error() != nil {
		log.Println("Failed to publish to topic")
		panic(token.Error())
	} else {
		log.Println("Published message")
	}
}

// this callback triggers when a message is received, it then prints the message (in the payload) and topic
var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	log.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

// upon connection to the client, this is called
var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	log.Println("Connected")
}

// this is called when the connection to the client is lost, it prints "Connection lost" and the corresponding error
var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	log.Printf("Connection lost: %v", err)
}
