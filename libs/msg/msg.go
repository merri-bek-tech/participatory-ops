package msg

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	paho "github.com/eclipse/paho.mqtt.golang"
)

type PahoConnection struct {
	DeviceId string
	Mqtt     paho.Client
}

func Connect(deviceId string) *PahoConnection {
	client := connectClient(
		"127.0.0.1",
		1883,
		deviceId,
	)

	return &PahoConnection{
		DeviceId: deviceId,
		Mqtt:     client,
	}
}

func (client *PahoConnection) GetGenericClient() MqttClient {
	return &PahoClient{
		Mqtt: &client.Mqtt,
	}
}

func (client *PahoConnection) GetMessenger() *Messenger {
	return &Messenger{
		DeviceId: client.DeviceId,
		client:   client.GetGenericClient(),
	}
}

func (client *PahoConnection) Disconnect() {
	client.Mqtt.Disconnect(250)
}

func (client *PahoConnection) PublishDetails(schemeId string, uuid string, details ComponentDetails) {
	text := encodeComponentDetails(details)
	client.Mqtt.Publish(deviceTopic(schemeId, uuid), AtMostOnce, false, text)
}

func (client *PahoConnection) SubscribeAllComponents(handlers CommsHandlers) {
	subscribe(allDevicesTopic(), client.Mqtt, handlers)
}

func (client *PahoConnection) SubscribeDevice(schemeId string, handlers CommsHandlers) {
	subscribe(deviceTopic(schemeId, client.DeviceId), client.Mqtt, handlers)
}

// PRIVATE

func connectClient(host string, port int, clientId string) paho.Client {
	opts := paho.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", host, port))
	opts.SetClientID(clientId) // set a name as you desire

	// create the client using the options above
	client := paho.NewClient(opts)
	// throw an error if the connection isn't successfull
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Println("Failed to connect to MQTT broker")
		log.Println(token.Error())
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

func handleComponentHeartbeatMessage(schemeId string, handlers CommsHandlers, _ Meta, contents string) {
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

	handlers.HandleHeartbeat(schemeId, heartbeat)
}

func handleDetailsRequestedMessage(schemeId string, handlers CommsHandlers, _ Meta, _ string) {
	if handlers.DetailsRequested == nil {
		return
	}

	handlers.DetailsRequested(schemeId)
}

func handleComponentDetailsMessage(schemeId string, handlers CommsHandlers, _ Meta, contents string) {
	if handlers.ComponentDetails == nil {
		return
	}

	var details ComponentDetails
	err := json.Unmarshal([]byte(contents), &details)
	if err != nil {
		log.Println("Failed to parse ComponentDetails: ", err)
		return
	}

	handlers.ComponentDetails(schemeId, details)
}

func handleParopsMessage(schemeId string, handlers CommsHandlers, meta Meta, contents string) {
	switch meta.Type {
	case "ComponentHeartbeat":
		handleComponentHeartbeatMessage(schemeId, handlers, meta, contents)
	case "DetailsRequested":
		handleDetailsRequestedMessage(schemeId, handlers, meta, contents)
	case "ComponentDetails":
		handleComponentDetailsMessage(schemeId, handlers, meta, contents)
	default:
		log.Printf("[%s] Unknown message type %s\n", schemeId, meta.Type)
	}
}

func handleMqttMessage(handlers CommsHandlers, message paho.Message) {
	//log.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())

	// split the message payload by |
	payload := string(message.Payload())
	payloadParts := strings.Split(payload, "|")

	schemeId := parseSchemeIdFromTopic(message.Topic())

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

	handleParopsMessage(schemeId, handlers, meta, contentsString)
}

func parseSchemeIdFromTopic(topic string) string {
	parts := strings.Split(topic, "/")

	if parts[0] != "schemes" {
		log.Println("Invalid topic format (expected schemes/xxx/...): ", topic)
	}

	return parts[1]
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
