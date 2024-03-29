package component_events

type ComponentHeartbeat struct {
	Uuid string `json:"uuid"`
	At   int64  `json:"at"`
}
