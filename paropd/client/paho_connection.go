package client

import (
	"fmt"
	"log"

	paho "github.com/eclipse/paho.mqtt.golang"
	msg "parops.libs/msg"
)

type PahoConnection struct {
	DeviceId string
	Mqtt     paho.Client
}

type MqttConnectionParams struct {
	Host string
	Port int
}

type OnConnectHandler func(*PahoConnection)

func Connect(deviceId string, connectionParams MqttConnectionParams, onConnect OnConnectHandler) *PahoConnection {
	client := connectClient(
		connectionParams.Host,
		connectionParams.Port,
		deviceId,
		func(client paho.Client) {
			onConnect(&PahoConnection{
				DeviceId: deviceId,
				Mqtt:     client,
			})
		},
	)

	return &PahoConnection{
		DeviceId: deviceId,
		Mqtt:     client,
	}
}

func (client *PahoConnection) GetGenericClient() msg.MqttClient {
	return &PahoClient{
		Mqtt: &client.Mqtt,
	}
}

func (client *PahoConnection) GetMessenger() *msg.Messenger {
	return &msg.Messenger{
		DeviceId: client.DeviceId,
		Client:   client.GetGenericClient(),
	}
}

func (client *PahoConnection) Disconnect() {
	client.Mqtt.Disconnect(250)
}

// PRIVATE

func connectClient(host string, port int, clientId string, onConnected paho.OnConnectHandler) paho.Client {
	opts := paho.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", host, port))
	opts.SetClientID(clientId)
	opts.SetAutoReconnect(true)
	opts.SetOnConnectHandler(onConnected)

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
