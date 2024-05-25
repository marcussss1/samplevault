package samplesv1

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c Controller) DislikeSoundByID(ctx echo.Context) error {
	userID := fmt.Sprint(ctx.Get("user_id"))
	soundID := ctx.Param("id")

	err := c.soundsService.DislikeSoundByID(ctx.Request().Context(), userID, soundID)
	if err != nil {
		return fmt.Errorf("dislike sound by id from sounds service: %w", err)
	}

	return ctx.NoContent(http.StatusOK)
}
