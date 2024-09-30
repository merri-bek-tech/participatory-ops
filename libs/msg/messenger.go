package msg

import (
	"encoding/json"
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
	topic := deviceTopic(schemeId, uuid)
	text := encodeDetailsRequested()

	log.Println("Publishing details requested", topic)

	messenger.client.Publish(topic, text)
}

func (messenger *Messenger) PublishDetails(schemeId string, uuid string, details ComponentDetails) {
	topic := deviceTopic(schemeId, uuid)
	text := encodeComponentDetails(details)
	messenger.client.Publish(topic, text)
}

func (messenger *Messenger) SubscribeDevice(schemeId string, handlers CommsHandlers) {
	topic := deviceTopic(schemeId, messenger.DeviceId)
	messenger.subscribe(topic, handlers)
}

func (messenger *Messenger) SubscribeAllComponents(handlers CommsHandlers) {
	topic := allDevicesTopic()
	messenger.subscribe(topic, handlers)
}

func (messenger *Messenger) subscribe(topic string, handlers CommsHandlers) {
	log.Println("Subscribing to topic: ", topic)

	messenger.client.Subscribe(topic, func(topic string, payload string) {
		// split the message payload by |
		payloadParts := strings.Split(payload, "|")

		schemeId := parseSchemeIdFromTopic(topic)

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
	})
}

// PRIVATE

func deviceTopic(schemeId string, uuid string) string {
	return "schemes/" + schemeId + "/components/" + uuid
}

func allDevicesTopic() string {
	return "schemes/+/components/+"
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

func jsonString(input any) string {
	output, err := json.Marshal(input)
	if err != nil {
		panic(err)
	}
	return string(output)
}

func parseSchemeIdFromTopic(topic string) string {
	parts := strings.Split(topic, "/")

	if parts[0] != "schemes" {
		log.Println("Invalid topic format (expected schemes/xxx/...): ", topic)
	}

	return parts[1]
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
	log.Println("Received details requested ")

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
