package msg

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	paho "github.com/eclipse/paho.mqtt.golang"
)

type Client struct {
	DeviceId string
	Mqtt     paho.Client
}

type CommsHandlers struct {
	HandleHeartbeat func(heartbeat ComponentHeartbeat)
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
		Mqtt:     client,
	}
}

func (client *Client) Disconnect() {
	client.Mqtt.Disconnect(250)
}

func (client *Client) PublishHeartbeat() {
	transmitHeartbeat("components/"+client.DeviceId, client.Mqtt, client.DeviceId)
}

func (client *Client) SubscribeAllComponents(handlers CommsHandlers) {
	subscribe("components/+", client.Mqtt, handlers)
}

func (client *Client) SubscribeDevice(handlers CommsHandlers) {
	subscribe("components/"+client.DeviceId, client.Mqtt, handlers)
}

// PRIVATE

func transmitHeartbeat(topic string, client paho.Client, clientId string) paho.Token {
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

func handleResult(token paho.Token) {
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
var messagePubHandler paho.MessageHandler = func(client paho.Client, msg paho.Message) {
	log.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

// upon connection to the client, this is called
var connectHandler paho.OnConnectHandler = func(client paho.Client) {
	log.Println("Connected")
}

// this is called when the connection to the client is lost, it prints "Connection lost" and the corresponding error
var connectLostHandler paho.ConnectionLostHandler = func(client paho.Client, err error) {
	log.Printf("Connection lost: %v", err)
}

func handleHeartbeatMessage(handlers CommsHandlers, _ Meta, contents string) {
	if handlers.HandleHeartbeat == nil {
		return
	}

	fmt.Println("Received heartbeat message: ", contents)

	var heartbeat ComponentHeartbeat
	err := json.Unmarshal([]byte(contents), &heartbeat)
	if err != nil {
		fmt.Println("Failed to parse heartbeat: ", err)
		return
	}

	handlers.HandleHeartbeat(heartbeat)
}

func handleParopsMessage(handlers CommsHandlers, meta Meta, contents string) {
	switch meta.Type {
	case "ComponentHeartbeat":
		handleHeartbeatMessage(handlers, meta, contents)
	default:
		fmt.Println("Unknown message type: ", meta.Type)
	}
}

func handleMqttMessage(handlers CommsHandlers, message paho.Message) {
	//fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())

	// split the message payload by |
	payload := string(message.Payload())
	payloadParts := strings.Split(payload, "|")

	if len(payloadParts) != 2 {
		fmt.Println("Invalid message format")
		return
	}

	// get the meta and the contents from the payloadParts
	metaString := payloadParts[0]
	contentsString := payloadParts[1]

	// parse Meta from the meta string
	var meta Meta
	err := json.Unmarshal([]byte(metaString), &meta)
	if err != nil {
		fmt.Println("Failed to parse meta")
		return
	}

	handleParopsMessage(handlers, meta, contentsString)
}

func subscribe(topic string, client paho.Client, handlers CommsHandlers) {
	handler := func(client paho.Client, message paho.Message) {
		handleMqttMessage(handlers, message)
	}

	// subscribe to the same topic, that was published to, to receive the messages
	token := client.Subscribe(topic, 1, handler)
	token.Wait()
	// Check for errors during subscribe (More on error reporting https://pkg.go.dev/github.com/eclipse/paho.mqtt.golang#readme-error-handling)
	if token.Error() != nil {
		fmt.Print("Failed to subscribe to topic\n")
		panic(token.Error())
	}
	fmt.Printf("Subscribed to topic: %s\n", topic)
}
