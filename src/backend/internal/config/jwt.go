package config

type JWTConfig struct {
	Secret string `yaml:"secret" env-required:"true"`
}
