package config

type ServiceConfig struct {
	Env  string `yaml:"env" env-required:"true"`
	Name string `yaml:"name" env-required:"true"`
}

const (
	EnvLocal = "local"
	EnvProd  = "prod"
)
