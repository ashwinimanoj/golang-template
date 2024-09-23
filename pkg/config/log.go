package config

type LoggerConfig struct {
	Level string `validate:"required"`
}
