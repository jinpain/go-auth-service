package config

type RedisConfig struct {
	Host     string `yaml:"host" env-required:"true"`
	Port     int    `yaml:"port" env-required:"true"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
	Protocol int    `yaml:"protocol" env-required:"true"`
}
