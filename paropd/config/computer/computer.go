package computer

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

type ComputedConfig struct {
	Uuid        string
	HostName    string
	ProductName string
	SysVendor   string
}

func ComputeConfig() *ComputedConfig {
	arch := archCommand()

	if arch == "x86_64" {
		return ComputeX86Config()
	} else if arch == "aarch64" {
		return ComputeArmConfig()
	} else {
		log.Println("Unknown architecture: ", arch)
		return &ComputedConfig{}
	}
}

// Private Compute Functions

func computeHostname() (result string, err error) {
	return os.Hostname()
}

// Private Helpers

func archCommand() string {
	return stringFromCommand("arch")
}

func stringFromCommand(command string) string {
	cmd := exec.Command(command)
	out, err := cmd.Output()
	if err != nil {
		// if there was any error, print it here
		fmt.Println("could not run command: ", command, err)
	}
	return strings.TrimSpace(string(out))
}

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
