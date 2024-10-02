package computer

func ComputeX86Config() *ComputedConfig {
	return &ComputedConfig{
		HostName:    orBlank(computeHostname()),
		ProductName: orBlank(computeDMIProductName()),
		SysVendor:   orBlank(computeDMISysVendor()),
	}
}

// Private Compute Functions

func computeDMIProductName() (result string, err error) {
	return stringFromFile("/sys/devices/virtual/dmi/id/product_name")
}

func computeDMISysVendor() (result string, err error) {
	return stringFromFile("/sys/devices/virtual/dmi/id/sys_vendor")
}
