package main

import (
	"github.com/rs/zerolog/log"
	"testtask/internal/pkg/app"
)

func main() {
	a, err := app.New()
	if err != nil {
		log.Fatal().Err(err).Msgf("Cannot create app")
	}

	err = a.Run()
	if err != nil {
		log.Fatal().Err(err).Msgf("Cannot run app")
	}
}
