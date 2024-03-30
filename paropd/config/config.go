package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"paropd/config/computed"
	"reflect"
	"strconv"

	"github.com/google/uuid"
)

type Config struct {
	Computed *computed.ComputedConfig
}

func LoadConfig(recompute bool) *Config {
	configDir := ensureBestConfigDir()

	log.Println("configDir:", configDir)

	computed := loadComputedConfig(configDir, recompute)

	return &Config{
		Computed: computed,
	}
}

// Private

func loadComputedConfig(configDir string, recompute bool) *computed.ComputedConfig {
	filePath := configDir + "/computed.pkl"

	// read existing computed config
	data := readExistingComputedConfig(filePath)

	// if it doesn't exist, or recompute is true, recompute it
	if data == nil || recompute {
		data = recomputeComputedConfig(data)

		// save it to disk
		writeComputedConfig(filePath, data)
	}

	return data
}

func readExistingComputedConfig(filePath string) *computed.ComputedConfig {
	if canReadFile(filePath) {
		computedConfig, err := computed.LoadFromPath(context.Background(), filePath)
		if err != nil {
			log.Println("Error loading computed config:", err)
			return nil
		}
		return computedConfig
	}
	return nil
}

func recomputeComputedConfig(existing *computed.ComputedConfig) *computed.ComputedConfig {
	changed := computed.ComputedConfig{
		HostName: hostnameOrDefault(""),
	}

	changed.Uuid = getUuid(existing)
	return &changed
}

func hostnameOrDefault(defaultHostname string) string {
	hostname, err := os.Hostname()
	if err != nil {
		return defaultHostname
	}
	return hostname
}

func getUuid(existing *computed.ComputedConfig) string {
	if existing != nil && existing.Uuid != "" {
		return existing.Uuid
	}
	return uuid.New().String()
}

func writeComputedConfig(filePath string, data *computed.ComputedConfig) {
	// open file for writing
	file, err := os.Create(filePath)
	if err != nil {
		log.Println("Error creating computed config file:", err)
		return
	}
	defer file.Close()

	// write data to file
	output := dummyPklOutput(data)
	_, err = file.WriteString(output)
	if err != nil {
		log.Println("Error writing computed config file:", err)
	}
}

// Pkl should do this for us, this is a hack
func dummyPklOutput(data *computed.ComputedConfig) string {
	s := *data
	v := reflect.ValueOf(s)
	typeOfS := v.Type()

	output := ""

	for i := 0; i < v.NumField(); i++ {
		key := typeOfS.Field(i).Tag.Get("pkl")
		value := dummyPklValueOutput(v.Field(i).Interface())
		line := fmt.Sprintf("%s = %s", key, value)
		output += line + "\n"
	}
	return output
}

func dummyPklValueOutput(input interface{}) string {
	if str, ok := input.(string); ok {
		return strconv.Quote(str)
	}

	return fmt.Sprintf("%v", input)
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
	err := os.Mkdir(dir, 0755)
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
