package config

type ServerConfig struct {
	Port int    `validate:"required"`
	Env  string `validate:"required"`
}
