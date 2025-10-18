package config

type Config struct {
	ServiceConfig           ServiceConfig  `yaml:"service"`
	DatabaseConfig          DatabaseConfig `yaml:"database"`
	RedisTokenConfig        RedisConfig    `yaml:"redis-token"`
	RedisVerificationConfig RedisConfig    `yaml:"redis-verification"`
	ServerConfig            ServerConfig   `yaml:"server"`
	SqlConfig               SqlConfig      `yaml:"sql"`
	JWTConfig               JWTConfig      `yaml:"jwt"`
}
