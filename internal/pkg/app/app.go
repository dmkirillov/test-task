package app

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"testtask/internal/app/endpoint"
	"testtask/internal/app/mv"
	"testtask/internal/service"
)

type App struct {
	e    *endpoint.Endpoint
	s    *service.Service
	echo *echo.Echo
}

func New() (*App, error) {
	a := &App{}

	a.s = service.New()

	a.e = endpoint.New(a.s)

	a.echo = echo.New()

	a.echo.Use(mv.RoleCheck)

	a.echo.GET("/status", a.e.Status)
	a.echo.GET("/health", a.e.Health)

	return a, nil
}

func (a *App) Run() error {
	log.Info().Msgf("Server has been started on port 8080")

	err := a.echo.Start(":8080")
	if err != nil {
		log.Fatal().Err(err).Msgf("Cannot start echo")
	}

	return nil
}
