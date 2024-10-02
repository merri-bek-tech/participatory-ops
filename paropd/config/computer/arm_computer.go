package computer

import (
	"fmt"
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
	return stringFromCpuInfo("Model"), nil
}

func computeSysVendor() (result string, err error) {
	model := stringFromCpuInfo("Model")
	if strings.Contains(model, "Raspberry Pi") {
		return "Raspberry Pi", nil
	} else {
		return "unknown", nil
	}
}

func stringFromCpuInfo(key string) (result string) {
	command := fmt.Sprintf("/proc/cpuinfo | grep '^%s'", key)
	line := stringFromCommand(command)
	value := strings.Split(line, ":")[1]
	return strings.TrimSpace(value)
}
