package discovery

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/Ullaakut/nmap/v3"
)

type BrokerLocation struct {
	Host string
	Port int
}

func FindBroker() *BrokerLocation {
	ports := "1883"
	targetIp := GetLocalIP()
	if targetIp == "" {
		targetIp = "192.168.1.0"
	}
	targetIp += "/24"

	log.Println("Scanning", targetIp, ports)

	hosts := getHosts(targetIp, ports)

	if len(hosts) == 0 {
		log.Println("No hosts found")
		return nil
	}

	// Output hosts
	for _, host := range hosts {
		log.Println("Host:", host.Addresses[0].Addr)
	}

	// By default, use the first host
	return brokerLocationFromHost(hosts[0])
}

func brokerLocationFromHost(host nmap.Host) *BrokerLocation {
	if len(host.Ports) == 0 || len(host.Addresses) == 0 {
		return nil
	}

	return &BrokerLocation{
		Host: host.Addresses[0].Addr,
		Port: int(host.Ports[0].ID),
	}
}

func getHosts(target string, ports string) []nmap.Host {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	// Equivalent to `/usr/local/bin/nmap -p 80,443,843 google.com facebook.com youtube.com`,
	// with a 5-minute timeout.
	scanner, err := nmap.NewScanner(
		ctx,
		nmap.WithTargets(target),
		nmap.WithPorts(ports),
		nmap.WithOpenOnly(),
		nmap.WithTimingTemplate(nmap.TimingAggressive),
	)
	if err != nil {
		log.Fatalf("unable to create nmap scanner: %v", err)
	}

	result, warnings, err := scanner.Run()
	if len(*warnings) > 0 {
		log.Printf("run finished with warnings: %s\n", *warnings) // Warnings are non-critical errors from nmap.
	}
	if err != nil {
		log.Fatalf("unable to run nmap scan: %v", err)
	}

	return result.Hosts
}

func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
