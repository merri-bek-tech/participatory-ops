package computer

import (
	"strings"
)

func ComputeArmConfig() *ComputedConfig {
	return &ComputedConfig{
		HostName:    orBlank(computeHostname()),
		ProductName: orBlank(computeProductName()),
		SysVendor:   orBlank(computeSysVendor()),
	}
}

// Private Compute Functions

func computeProductName() (result string, err error) {
	return stringFromCpuInfo("Model")
}

func computeSysVendor() (result string, err error) {
	model, err := stringFromCpuInfo("Model")

	if err != nil {
		return "", err
	}

	if strings.Contains(model, "Raspberry Pi") {
		return "Raspberry Pi", nil
	} else {
		return "unknown", nil
	}
}

func stringFromCpuInfo(key string) (result string, err error) {
	line, fileErr := stringFromFileStartingWith("/proc/cpuinfo", key)
	if fileErr != nil {
		return "", fileErr
	}

	value := strings.Split(line, ":")[1]
	return strings.TrimSpace(value), nil
}
