package msg

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	paho "github.com/eclipse/paho.mqtt.golang"
)

var (
	AtMostOnce  byte = 0
	AtLeastOnce byte = 1
	ExactlyOnce byte = 2
)

type Client struct {
	DeviceId string
	Mqtt     paho.Client
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

func (client *Client) PublishMyHeartbeat() {
	payload := ComponentHeartbeat{
		Uuid: client.DeviceId,
		At:   time.Now().Unix(),
	}

	client.publishHeartbeat(deviceTopic(client.DeviceId), payload)
}

func (client *Client) PublishDetailsRequested(uuid string) {
	client.publishDetailsRequested(deviceTopic(uuid))
}

func (client *Client) PublishDetails(uuid string, details ComponentDetails) {
	text := encodeComponentDetails(details)
	client.Mqtt.Publish(deviceTopic(uuid), AtMostOnce, false, text)
}

func (client *Client) SubscribeAllComponents(handlers CommsHandlers) {
	subscribe(allDevicesTopic(), client.Mqtt, handlers)
}

func (client *Client) SubscribeDevice(handlers CommsHandlers) {
	subscribe(deviceTopic(client.DeviceId), client.Mqtt, handlers)
}

// PRIVATE

func deviceTopic(uuid string) string {
	return "components/" + uuid
}

func allDevicesTopic() string {
	return "components/+"
}

func (client *Client) publishHeartbeat(topic string, data ComponentHeartbeat) paho.Token {
	log.Println("Publishing heartbeat")

	text := encodeHeartbeat(data)
	return client.Mqtt.Publish(topic, AtMostOnce, false, text)
}

func (client *Client) publishDetailsRequested(topic string) paho.Token {
	log.Println("Publishing details requested")

	text := encodeDetailsRequested()
	return client.Mqtt.Publish(topic, AtMostOnce, false, text)
}

func encodeHeartbeat(heartbeat ComponentHeartbeat) string {
	meta := Meta{
		Type:    "ComponentHeartbeat",
		Version: "1.0",
	}

	return encodeMessage(meta, heartbeat)
}

func encodeDetailsRequested() string {
	meta := Meta{
		Type:    "DetailsRequested",
		Version: "1.0",
	}

	return encodeMessage(meta, nil)
}

func encodeComponentDetails(details ComponentDetails) string {
	meta := Meta{
		Type:    "ComponentDetails",
		Version: "1.0",
	}

	return encodeMessage(meta, details)
}

func encodeMessage(meta Meta, body any) string {
	metaText := jsonString(meta)
	payloadText := ""
	if body != nil {
		payloadText = jsonString(body)
	}

	return strings.Join([]string{metaText, payloadText}, "|")
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

func handleComponentHeartbeatMessage(handlers CommsHandlers, _ Meta, contents string) {
	if handlers.HandleHeartbeat == nil {
		return
	}

	log.Println("Received heartbeat message: ", contents)

	var heartbeat ComponentHeartbeat
	err := json.Unmarshal([]byte(contents), &heartbeat)
	if err != nil {
		log.Println("Failed to parse heartbeat: ", err)
		return
	}

	handlers.HandleHeartbeat(heartbeat)
}

func handleDetailsRequestedMessage(handlers CommsHandlers, _ Meta, _ string) {
	if handlers.DetailsRequested == nil {
		return
	}

	handlers.DetailsRequested()
}

func handleComponentDetailsMessage(handlers CommsHandlers, _ Meta, contents string) {
	if handlers.ComponentDetails == nil {
		return
	}

	var details ComponentDetails
	err := json.Unmarshal([]byte(contents), &details)
	if err != nil {
		log.Println("Failed to parse ComponentDetails: ", err)
		return
	}

	handlers.ComponentDetails(details)
}

func handleParopsMessage(handlers CommsHandlers, meta Meta, contents string) {
	switch meta.Type {
	case "ComponentHeartbeat":
		handleComponentHeartbeatMessage(handlers, meta, contents)
	case "DetailsRequested":
		handleDetailsRequestedMessage(handlers, meta, contents)
	case "ComponentDetails":
		handleComponentDetailsMessage(handlers, meta, contents)
	default:
		log.Println("Unknown message type: ", meta.Type)
	}
}

func handleMqttMessage(handlers CommsHandlers, message paho.Message) {
	//log.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())

	// split the message payload by |
	payload := string(message.Payload())
	payloadParts := strings.Split(payload, "|")

	if len(payloadParts) != 2 {
		log.Println("Invalid message format")
		return
	}

	// get the meta and the contents from the payloadParts
	metaString := payloadParts[0]
	contentsString := payloadParts[1]

	// parse Meta from the meta string
	var meta Meta
	err := json.Unmarshal([]byte(metaString), &meta)
	if err != nil {
		log.Println("Failed to parse meta")
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
		log.Print("Failed to subscribe to topic\n")
		panic(token.Error())
	}
	log.Printf("Subscribed to topic: %s\n", topic)
}
