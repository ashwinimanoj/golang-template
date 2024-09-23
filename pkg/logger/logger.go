package logger

import (
	"github.com/ashwinimanoj/golang-template/pkg/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

const defaultLogLevel = zerolog.TraceLevel

func InitLogger(config *config.Config) {
	zerolog.SetGlobalLevel(zerolog.Level(getLogLevel(config.Logger.Level)))
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
}

func getLogLevel(level string) zerolog.Level {
	logLevel, err := zerolog.ParseLevel(level)
	if err != nil {
		log.Warn().Msgf("invalid loglevel: %s, replaced with level: %s", level, defaultLogLevel)
		return defaultLogLevel
	}
	return logLevel
}
