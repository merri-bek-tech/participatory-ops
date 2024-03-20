package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
)

// this callback triggers when a message is received, it then prints the message (in the payload) and topic
var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

// upon connection to the client, this is called
var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

// this is called when the connection to the client is lost, it prints "Connection lost" and the corresponding error
var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connection lost: %v", err)
}

func main() {
	clientId := uuid.New().String()

	client := connectClient(
		"82e12caef57c4c8288d08fe23854c097.s1.eu.hivemq.cloud",
		8883,
		"paropd",
		"be9eiQuo",
		clientId,
	)

	// subscribe(client)
	publishLoop(client, clientId)

	client.Disconnect(250)
}

func connectClient(host string, port int, username string, password string, clientId string) mqtt.Client {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tls://%s:%d", host, port))
	opts.SetClientID(clientId) // set a name as you desire
	opts.SetUsername(username) // these are the credentials that you declare for your cluster (see readme)
	opts.SetPassword(password)
	// (optionally) configure callback handlers that get called on certain events
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	// create the client using the options above
	client := mqtt.NewClient(opts)
	// throw an error if the connection isn't successfull
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	return client
}

// func subscribe(client mqtt.Client) {
// 	// subscribe to the same topic, that was published to, to receive the messages
// 	topic := "topic/test"
// 	token := client.Subscribe(topic, 1, nil)
// 	token.Wait()
// 	// Check for errors during subscribe (More on error reporting https://pkg.go.dev/github.com/eclipse/paho.mqtt.golang#readme-error-handling)
// 	if token.Error() != nil {
// 		fmt.Print("Failed to subscribe to topic\n")
// 		panic(token.Error())
// 	}
// 	fmt.Printf("Subscribed to topic: %s\n", topic)
// }

func publishLoop(client mqtt.Client, clientId string) {
	for i := 0; i < 2; i++ {
		token := publishHeartbeat(client, clientId)
		handleResult(token)
		time.Sleep(time.Second)
	}
}

func handleResult(token mqtt.Token) {
	token.Wait()
	// Check for errors during publishing (More on error reporting https://pkg.go.dev/github.com/eclipse/paho.mqtt.golang#readme-error-handling)
	if token.Error() != nil {
		fmt.Printf("Failed to publish to topic")
		panic(token.Error())
	} else {
		fmt.Printf("Published message\n")
	}
}

type ComponentHeartbeat struct {
	Uuid string `json:"uuid"`
	At   int64  `json:"at"`
}

type Meta struct {
	Type    string `json:"type"`
	Version string `json:"version"`
}

func jsonString(input any) string {
	output, err := json.Marshal(input)
	if err != nil {
		panic(err)
	}
	return string(output)
}

func publishHeartbeat(client mqtt.Client, clientId string) mqtt.Token {
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

	return client.Publish("topic/test", qos, false, strings.Join([]string{metaText, payloadText}, "|"))
}
