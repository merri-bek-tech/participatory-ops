package config

import (
	"fmt"
	"os"
	"paropd/config/computed"
)

type Config struct {
	Computed *computed.ComputedConfig
}

func LoadConfig(recompute bool) (*Config, error) {
	configDir := ensureBestConfigDir()

	fmt.Println("configDir:", configDir)

	// file := configDir + "/computed.pkl"
	// info, err := os.Stat(file)
	// if err != nil {
	// 	fmt.Printf("error getting file info: %v\n", err)
	// 	return nil, err
	// }
	// if info.Mode().Perm()&0444 == 0444 {
	// 	fmt.Println("file has readable permission")
	// } else {
	// 	fmt.Println("file does not have readable permission")
	// }

	// if recompute {
	// 	return nil, nil
	// }

	// computedConfig, err := computed.LoadFromPath(context.Background(), configDir+"/computed.pkl")
	// if err != nil {
	// 	return nil, err
	// }

	// return &Config{Computed: computedConfig}, nil
	return nil, nil
}

// Private

func ensureBestConfigDir() string {
	basename := "paropd"
	candidateDirs := []string{"/etc/" + basename, "~/.config/" + basename}

	for _, dir := range candidateDirs {
		if canReadOrMake(dir) {
			return dir
		}
	}

	panic("Could not find a suitable config directory")
}

func canReadOrMake(dir string) bool {
	absDir := replaceTildeWithHomeDir(dir)
	_, err := os.Stat(absDir)

	if err == nil {
		return true
	} else if os.IsNotExist(err) {
		return canMake(absDir)
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
