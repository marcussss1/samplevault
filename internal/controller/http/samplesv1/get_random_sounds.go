package samplesv1

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (c Controller) GetRandomSounds(ctx echo.Context) error {
	// TODO Временно, когда появится авторизация нужно поменять логику
	sounds, err := c.soundsService.GetRandomSounds(ctx.Request().Context())
	if err != nil {
		return fmt.Errorf("get random sounds from sounds service: %w", err)
	}

	return ctx.JSON(http.StatusOK, sounds)
}
