package component_comms

import (
	"encoding/json"
	"fmt"
	"strings"

	"parops.libs/msg"

	paho "github.com/eclipse/paho.mqtt.golang"
	uuid "github.com/google/uuid"
)

type CommsHandlers struct {
	HandleHeartbeat func(heartbeat msg.ComponentHeartbeat)
}

func MonitorComponents(handlers CommsHandlers) {
	deviceId := "api-" + uuid.New().String()
	client := msg.Connect(deviceId)

	subscribe("components/+", client.Mqtt, handlers)
}

// PRIVATE

func handleHeartbeatMessage(handlers CommsHandlers, _ msg.Meta, contents string) {
	fmt.Println("Received heartbeat message: ", contents)

	var heartbeat msg.ComponentHeartbeat
	err := json.Unmarshal([]byte(contents), &heartbeat)
	if err != nil {
		fmt.Println("Failed to parse heartbeat: ", err)
		return
	}

	handlers.HandleHeartbeat(heartbeat)
}

func handleParopsMessage(handlers CommsHandlers, meta msg.Meta, contents string) {
	switch meta.Type {
	case "ComponentHeartbeat":
		handleHeartbeatMessage(handlers, meta, contents)
	default:
		fmt.Println("Unknown message type: ", meta.Type)
	}
}

func handleMqttMessage(handlers CommsHandlers, client paho.Client, message paho.Message) {
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
	var meta msg.Meta
	err := json.Unmarshal([]byte(metaString), &meta)
	if err != nil {
		fmt.Println("Failed to parse meta")
		return
	}

	handleParopsMessage(handlers, meta, contentsString)
}

func subscribe(topic string, client paho.Client, handlers CommsHandlers) {
	handler := func(client paho.Client, msg paho.Message) {
		handleMqttMessage(handlers, client, msg)
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
