package config

type RedisConfig struct {
	Host     string `yaml:"host" env:"HOST" env-required:"true"`
	Port     int    `yaml:"port" env:"PORT" env-required:"true"`
	Password string `yaml:"password" env:"PASSWORD"`
	DB       int    `yaml:"db" env:"DB"`
	Protocol int    `yaml:"protocol" env:"PROTOCOL" env-required:"true"`
}
