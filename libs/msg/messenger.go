package msg

import (
	"log"
	"time"
)

type Messenger struct {
	DeviceId string
	client   MqttClient
}

func (messenger *Messenger) PublishMyHeartbeat(schemeId string) {
	messenger.PublishHeartbeat(schemeId, messenger.DeviceId)
}

func (messenger *Messenger) PublishHeartbeat(schemeId string, deviceId string) {
	payload := ComponentHeartbeat{
		Uuid: deviceId,
		At:   time.Now().Unix(),
	}

	log.Println("Publishing heartbeat")

	topic := deviceTopic(schemeId, deviceId)
	text := encodeHeartbeat(payload)
	messenger.client.Publish(topic, text)
}
