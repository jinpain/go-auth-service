package config

type Config struct {
	ServiceConfig           ServiceConfig  `yaml:"service"`
	DatabaseConfig          DatabaseConfig `yaml:"database"`
	RedisTokenConfig        RedisConfig    `yaml:"redis-token" envPrefix:"REDIS_TOKEN_"`
	RedisVerificationConfig RedisConfig    `yaml:"redis-verification" envPrefix:"REDIS_VERIFICATION_"`
	ServerConfig            ServerConfig   `yaml:"server"`
	SqlConfig               SqlConfig      `yaml:"sql"`
	JWTConfig               JWTConfig      `yaml:"jwt"`
}
