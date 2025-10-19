package config

type RedisConfig struct {
	Host     string `yaml:"host" env:"HOST"`
	Port     int    `yaml:"port" env:"PORT"`
	Password string `yaml:"password" env:"PASSWORD"`
	DB       int    `yaml:"db" env:"DB"`
	Protocol int    `yaml:"protocol" env:"PROTOCOL"`
}
