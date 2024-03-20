package main

import (
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
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
	var broker = "82e12caef57c4c8288d08fe23854c097.s1.eu.hivemq.cloud" // find the host name in the Overview of your cluster (see readme)
	var port = 8883                                                    // find the port right under the host name, standard is 8883
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tls://%s:%d", broker, port))
	opts.SetClientID("test1")  // set a name as you desire
	opts.SetUsername("paropd") // these are the credentials that you declare for your cluster (see readme)
	opts.SetPassword("be9eiQuo")
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

	// subscribe(client)
	publish(client)

	client.Disconnect(250)
}

func subscribe(client mqtt.Client) {
	// subscribe to the same topic, that was published to, to receive the messages
	topic := "topic/test"
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	// Check for errors during subscribe (More on error reporting https://pkg.go.dev/github.com/eclipse/paho.mqtt.golang#readme-error-handling)
	if token.Error() != nil {
		fmt.Print("Failed to subscribe to topic\n")
		panic(token.Error())
	}
	fmt.Printf("Subscribed to topic: %s\n", topic)
}

func publish(client mqtt.Client) {
	// publish the message "Message" to the topic "topic/test" 10 times in a for loop
	num := 10
	for i := 0; i < num; i++ {
		text := fmt.Sprintf("Message %d", i)
		token := client.Publish("topic/test", 0, false, text)
		token.Wait()
		// Check for errors during publishing (More on error reporting https://pkg.go.dev/github.com/eclipse/paho.mqtt.golang#readme-error-handling)
		if token.Error() != nil {
			fmt.Printf("Failed to publish to topic")
			panic(token.Error())
		} else {
			fmt.Printf("Published message %d\n", i)
		}
		time.Sleep(time.Second)
	}
}
