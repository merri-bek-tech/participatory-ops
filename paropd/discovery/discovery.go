package discovery

type BrokerLocation struct {
	Host string
	Port int
}

func FindBroker() *BrokerLocation {
	return &BrokerLocation{
		Host: "127.0.0.1",
		Port: 1883,
	}
}
