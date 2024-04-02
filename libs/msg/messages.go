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
	Uuid        string `json:"uuid"`
	HostName    string `json:"hostName"`
	ProductName string `json:"productName"`
	SysVendor   string `json:"sysVendor"`
}

type CommsHandlers struct {
	HandleHeartbeat  func(schemeId string, heartbeat ComponentHeartbeat)
	DetailsRequested func(schemeId string)
	ComponentDetails func(schemeId string, details ComponentDetails)
}
