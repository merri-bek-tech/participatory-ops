package config

import (
	"context"
	"fmt"
	"os"
	"paropd/config/computed"
)

type Config struct {
	Computed *computed.ComputedConfig
}

func LoadConfig(recompute bool) *Config {
	configDir := ensureBestConfigDir()

	fmt.Println("configDir:", configDir)

	computed := loadComputedConfig(configDir, recompute)

	return &Config{
		Computed: computed,
	}
}

// Private

func loadComputedConfig(configDir string, recompute bool) *computed.ComputedConfig {
	filePath := configDir + "/computed.pkl"

	if canReadFile(filePath) {
		computedConfig, err := computed.LoadFromPath(context.Background(), configDir+"/computed.pkl")
		if err != nil {
			fmt.Println("Error loading computed config:", err)
			return &computed.ComputedConfig{}
		}
		return computedConfig
	}

	return &computed.ComputedConfig{}
}

func candidateConfigDirs() []string {
	basename := "paropd"

	return []string{
		"/etc/" + basename,
		replaceTildeWithHomeDir("~/.config/" + basename),
	}
}

func ensureBestConfigDir() string {
	candidateDirs := candidateConfigDirs()

	for _, dir := range candidateDirs {
		if canReadOrMakeDir(dir) {
			return dir
		}
	}

	panic("Could not find a suitable config directory")
}

func canReadFile(path string) bool {
	_, err := os.Stat(path)

	if err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		fmt.Println("Error checking for file:", err)
	}
	return false
}

func canReadOrMakeDir(dir string) bool {
	_, err := os.Stat(dir)

	if err == nil {
		return true
	} else if os.IsNotExist(err) {
		return canMake(dir)
	} else {
		fmt.Println("Error checking for directory:", err)
	}
	return false
}

func canMake(dir string) bool {
	err := os.Mkdir(dir, 0755)
	if err != nil {
		fmt.Println("Error making directory:", err)
		return false
	}
	return true
}

func replaceTildeWithHomeDir(dir string) string {
	if dir[0] == '~' {
		homedir, err := os.UserHomeDir()
		if err != nil {
			panic(err)
		}
		return homedir + dir[1:]
	}
	return dir
}
