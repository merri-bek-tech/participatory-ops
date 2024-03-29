package config

import (
	"context"
	"paropd/config/computed"
)

type Config struct {
	Computed *computed.ComputedConfig
}

func LoadConfig(recompute bool) (*Config, error) {
	if recompute {
		return nil, nil
	}

	computedConfig, err := computed.LoadFromPath(context.Background(), "./config/computed/sample.pkl")
	if err != nil {
		return nil, err
	}

	return &Config{Computed: computedConfig}, nil
}
