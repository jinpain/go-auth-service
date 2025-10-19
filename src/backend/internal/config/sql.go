package config

type SqlConfig struct {
	Path string `yaml:"path" env:"SQL_PATH"`
}
