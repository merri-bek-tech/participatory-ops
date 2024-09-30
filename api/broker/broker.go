package broker

import (
	"log"

	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/hooks/auth"
	"github.com/mochi-mqtt/server/v2/listeners"
)

func MessageBroker(onStarted func(inlineClient *InlineClient)) {
	log.Println("Starting MQTT broker")

	started := make(chan bool, 1)

	// Create the new MQTT Server.
	server := mqtt.New(&mqtt.Options{
		InlineClient: true,
	})

	// Allow all connections.
	_ = server.AddHook(new(auth.AllowHook), nil)

	// use the onStarted hook
	err := server.AddHook(&LifecycleHook{}, &LifecycleHookOptions{started: started})
	if err != nil {
		log.Fatal(err)
	}

	// Create a TCP listener on a standard port.
	tcp := listeners.NewTCP(listeners.Config{ID: "t1", Address: ":1883"})
	err = server.AddListener(tcp)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		err := server.Serve()
		if err != nil {
			log.Fatal(err)
		}
	}()

	<-started

	if onStarted != nil {
		server.Publish("direct/publish", []byte("packet scheduled message"), false, 0)
		inlineClient := BuildInlineClient(server)

		onStarted(inlineClient)
	}
}
