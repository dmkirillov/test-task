package endpoint

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Endpoint struct {
	s Service
}

type Service interface {
	DaysLeft() int64
}

func New(s Service) *Endpoint {
	return &Endpoint{s: s}
}

func (e *Endpoint) Status(ctx echo.Context) error {
	s := fmt.Sprintf("Days until 1st January 2024 is %v", e.s.DaysLeft())

	err := ctx.String(http.StatusOK, s)
	if err != nil {
		return err
	}
	return nil
}

func (e *Endpoint) Health(ctx echo.Context) error {
	err := ctx.String(http.StatusOK, "OK.")

	if err != nil {
		return err
	}
	return nil
}
