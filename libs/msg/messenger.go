package msg

import (
	"log"
	"strings"
	"time"
)

type Messenger struct {
	DeviceId string
	client   MqttClient
}

func (messenger *Messenger) PublishMyHeartbeat(schemeId string) {
	payload := ComponentHeartbeat{
		Uuid: messenger.DeviceId,
		At:   time.Now().Unix(),
	}

	log.Println("Publishing heartbeat")

	topic := deviceTopic(schemeId, messenger.DeviceId)
	text := encodeHeartbeat(payload)
	messenger.client.Publish(topic, text)
}

func (messenger *Messenger) PublishDetailsRequested(schemeId string, uuid string) {
	topic := deviceTopic(schemeId, messenger.DeviceId)
	text := encodeDetailsRequested()
	messenger.client.Publish(topic, text)
}

func (messenger *Messenger) PublishDetails(schemeId string, uuid string, details ComponentDetails) {
	topic := deviceTopic(schemeId, uuid)
	text := encodeComponentDetails(details)
	messenger.client.Publish(topic, text)
}

// PRIVATE

func deviceTopic(schemeId string, uuid string) string {
	return "schemes/" + schemeId + "/components/" + uuid
}

func encodeHeartbeat(heartbeat ComponentHeartbeat) string {
	meta := Meta{
		Type:    "ComponentHeartbeat",
		Version: "1.0",
	}

	return encodeMessage(meta, heartbeat)
}

func encodeComponentDetails(details ComponentDetails) string {
	meta := Meta{
		Type:    "ComponentDetails",
		Version: "1.0",
	}

	return encodeMessage(meta, details)
}

func allDevicesTopic() string {
	return "schemes/+/components/+"
}

func encodeDetailsRequested() string {
	meta := Meta{
		Type:    "DetailsRequested",
		Version: "1.0",
	}

	return encodeMessage(meta, nil)
}

func encodeMessage(meta Meta, body any) string {
	metaText := jsonString(meta)
	payloadText := ""
	if body != nil {
		payloadText = jsonString(body)
	}

	return strings.Join([]string{metaText, payloadText}, "|")
}
