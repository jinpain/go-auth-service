package configloader

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

const (
	EnvLocal = "local"
	EnvProd  = "prod"
)

func Load[ConfigType any]() (*ConfigType, error) {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		return nil, fmt.Errorf("environment variable CONFIG_PATH is empty")
	}

	return LoadPath[ConfigType](configPath)
}

func LoadPath[ConfigType any](configPath string) (*ConfigType, error) {
	info, err := os.Stat(configPath)
	if err != nil {
		return nil, fmt.Errorf("cannot stat config file %q: %w", configPath, err)
	}
	if info.IsDir() {
		return nil, fmt.Errorf("config path %q is a directory, expected a file", configPath)
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("cannot read config file %q: %w", configPath, err)
	}

	var cfg ConfigType
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("cannot unmarshal yaml config: %w", err)
	}

	return &cfg, nil
}
