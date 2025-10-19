package configloader

import (
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
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
	var cfg ConfigType

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return nil, fmt.Errorf("cannot read config file %q: %w", configPath, err)
	}

	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, fmt.Errorf("cannot read environment variables: %w", err)
	}

	return &cfg, nil
}
