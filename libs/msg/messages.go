package msg

type Meta struct {
	Type    string `json:"type"`
	Version string `json:"version"`
}

type ComponentHeartbeat struct {
	Uuid string `json:"uuid"`
	At   int64  `json:"at"`
}

type ComponentDetails struct {
	Uuid     string `json:"uuid"`
	HostName string `json:"hostName"`
}

type CommsHandlers struct {
	HandleHeartbeat  func(heartbeat ComponentHeartbeat)
	DetailsRequested func()
	ComponentDetails func(details ComponentDetails)
}