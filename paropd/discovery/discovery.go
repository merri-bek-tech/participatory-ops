package discovery

type BrokerLocation struct {
	Host string
	Port int
}

func FindBroker() *BrokerLocation {
	// return nil
	return &BrokerLocation{
		Host: "127.0.0.1",
		Port: 1883,
	}
}
