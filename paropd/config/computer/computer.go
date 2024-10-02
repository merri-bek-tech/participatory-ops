package computer

import (
	"log"
	"os"
	"strings"
)

type ComputedConfig struct {
	Uuid        string
	HostName    string
	ProductName string
	SysVendor   string
}

func ComputeConfig() *ComputedConfig {
	return &ComputedConfig{
		HostName:    orBlank(computeHostname()),
		ProductName: orBlank(computeProductName()),
		SysVendor:   orBlank(computeSysVendor()),
	}
}

// Private Compute Functions

func computeHostname() (result string, err error) {
	return os.Hostname()
}

func computeProductName() (result string, err error) {
	return stringFromFile("/sys/devices/virtual/dmi/id/product_name")
}

func computeSysVendor() (result string, err error) {
	return stringFromFile("/sys/devices/virtual/dmi/id/sys_vendor")
}

// Private Helpers

func orBlank(result string, err error) string {
	if err != nil {
		return ""
	}
	return result
}

func stringFromFile(filePath string) (result string, err error) {
	bytes, err := os.ReadFile(filePath) // just pass the file name
	if err != nil {
		log.Println("Error reading file: ", err)
	}

	str := strings.TrimSpace(string(bytes))
	return str, nil
}
