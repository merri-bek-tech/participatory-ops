package component_comms

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	uuid "github.com/google/uuid"
)

type CommsHandlers struct {
	HandleHeartbeat func(heartbeat ComponentHeartbeat)
}

func MonitorComponents(handlers CommsHandlers) {
	clientId := "api-" + uuid.New().String()

	client := connectClient(
		"82e12caef57c4c8288d08fe23854c097.s1.eu.hivemq.cloud",
		8883,
		"paropd",
		"be9eiQuo",
		clientId,
	)

	subscribe(client, handlers)

	fmt.Println("Monitoring components finished")
}

type Meta struct {
	Type    string `json:"type"`
	Version string `json:"version"`
}

type ComponentHeartbeat struct {
	Uuid string `json:"uuid"`
	At   int64  `json:"at"`
}

// PRIVATE

func handleHeartbeatMessage(handlers CommsHandlers, _ Meta, contents string) {
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

func handleMqttMessage(handlers CommsHandlers, client mqtt.Client, msg mqtt.Message) {
	//fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())

	// split the message payload by |
	payload := string(msg.Payload())
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

func subscribe(client mqtt.Client, handlers CommsHandlers) {
	handler := func(client mqtt.Client, msg mqtt.Message) {
		handleMqttMessage(handlers, client, msg)
	}

	// subscribe to the same topic, that was published to, to receive the messages
	topic := "topic/test"
	token := client.Subscribe(topic, 1, handler)
	token.Wait()
	// Check for errors during subscribe (More on error reporting https://pkg.go.dev/github.com/eclipse/paho.mqtt.golang#readme-error-handling)
	if token.Error() != nil {
		fmt.Print("Failed to subscribe to topic\n")
		panic(token.Error())
	}
	fmt.Printf("Subscribed to topic: %s\n", topic)
}

func connectClient(host string, port int, username string, password string, clientId string) mqtt.Client {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tls://%s:%d", host, port))
	opts.SetClientID(clientId) // set a name as you desire
	opts.SetUsername(username) // these are the credentials that you declare for your cluster (see readme)
	opts.SetPassword(password)
	opts.SetTLSConfig(&tls.Config{InsecureSkipVerify: true})

	// (optionally) configure callback handlers that get called on certain events
	// opts.SetDefaultPublishHandler(messagePubHandler)
	// opts.OnConnect = connectHandler
	// opts.OnConnectionLost = connectLostHandler
	// create the client using the options above
	client := mqtt.NewClient(opts)
	// throw an error if the connection isn't successfull
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	return client
}
