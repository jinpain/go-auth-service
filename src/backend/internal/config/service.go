package config

type ServiceConfig struct {
	Env  string `yaml:"env" env:"SERVICE_ENV"`
	Name string `yaml:"name" env:"SERVICE_NAME"`
}
