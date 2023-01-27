package mv

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func RoleCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		userRole := ctx.Request().Header.Get("User-Role")

		if userRole == "admin" {
			log.Info().Msgf("red button user detected")
		}

		err := next(ctx)
		if err != nil {
			return err
		}

		return nil
	}
}
