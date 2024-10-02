package config

import (
	"log"
	"os"

	"paropd/config/computer"

	"github.com/BurntSushi/toml"
	"github.com/google/uuid"
)

type Config struct {
	Computed *computer.ComputedConfig
	SchemeId string
}

func LoadConfig(recompute bool) *Config {
	configDir := ensureBestConfigDir()

	log.Println("configDir:", configDir)

	computed := loadComputedConfig(configDir, recompute)

	return &Config{
		Computed: computed,
		SchemeId: "mbt-dev",
	}
}

// Private

func loadComputedConfig(configDir string, recompute bool) *computer.ComputedConfig {
	filePath := configDir + "/computed.toml"

	// read existing computed config
	data := readExistingComputedConfig(filePath)

	// if it doesn't exist, or recompute is true, recompute it
	if data == nil || recompute {
		data = recomputeComputedConfig(data)

		// 	save it to disk
		writeComputedConfig(filePath, data)
	}

	return data
}

func readExistingComputedConfig(filePath string) *computer.ComputedConfig {
	if canReadFile(filePath) {
		var computedConfig *computer.ComputedConfig
		_, err := toml.DecodeFile(filePath, &computedConfig)
		if err != nil {
			log.Println("Error loading computed config:", err)
			return nil
		}
		return computedConfig
	}
	return nil
}

func recomputeComputedConfig(existing *computer.ComputedConfig) *computer.ComputedConfig {
	changed := computer.ComputeConfig()

	changed.Uuid = getUuid(existing)
	return changed
}

func getUuid(existing *computer.ComputedConfig) string {
	if existing != nil && existing.Uuid != "" {
		return existing.Uuid
	}
	return uuid.New().String()
}

func writeComputedConfig(filePath string, data *computer.ComputedConfig) {
	// open file for writing
	file, err := os.Create(filePath)
	if err != nil {
		log.Println("Error creating computed config file:", err)
		return
	}
	defer file.Close()

	encoder := toml.NewEncoder(file)

	// write data to file
	err = encoder.Encode(data)
	if err != nil {
		log.Println("Error writing computed config file:", err)
	}
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
		log.Println("Error checking for file:", err)
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
		log.Println("Error checking for directory:", err)
	}
	return false
}

func canMake(dir string) bool {
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		log.Println("Error making directory:", err)
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
