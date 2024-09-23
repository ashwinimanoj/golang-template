package main

import (
	"github.com/ashwinimanoj/golang-template/pkg/config"
	"github.com/ashwinimanoj/golang-template/pkg/logger"
	"github.com/rs/zerolog/log"
)

func main() {
	config, errs := config.NewConfig()
	if errs != nil {
		log.Fatal().Errs("errors", errs).Msg("failed to initialize config")
	}

	log.Trace().Msgf("with config: %+v", config)
	logger.InitLogger(config)

	log.Info().Msg("starting server")
}
