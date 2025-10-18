package config

type ServiceConfig struct {
	Env  string `yaml:"env" env:"SERVICE_ENV" env-required:"true"`
	Name string `yaml:"name" env:"SERVICE_NAME" env-required:"true"`
}
