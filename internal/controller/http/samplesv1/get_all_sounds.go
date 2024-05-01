package samplesv1

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c Controller) GetAllSounds(ctx echo.Context) error {
	sounds, err := c.soundsService.GetAllSounds(ctx.Request().Context())
	if err != nil {
		return fmt.Errorf("get all sounds from sounds service: %w", err)
	}

	fmt.Println("sounds: ", sounds)

	return ctx.JSON(http.StatusOK, sounds)
}
