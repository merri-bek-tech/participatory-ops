package broker

import (
	"log"

	"github.com/google/uuid"
	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/hooks/auth"
	"github.com/mochi-mqtt/server/v2/listeners"
	msg "parops.libs/msg"
)

func MessageBroker(onStarted func(messenger *msg.Messenger)) {
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
		// This is temporary, we need this to stay constant
		deviceId := "api-" + uuid.New().String()

		messenger := &msg.Messenger{
			DeviceId: deviceId,
			Client:   BuildInlineClient(server),
		}

		onStarted(messenger)
	}
}
