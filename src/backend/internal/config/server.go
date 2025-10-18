package config

type ServerConfig struct {
	Host string `yaml:"host" env:"SERVER_HOST" env-required:"true"`
	Port int    `yaml:"port" env:"SERVER_PORT" env-required:"true"`
}
