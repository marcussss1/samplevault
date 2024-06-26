package samplesv1

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c Controller) LikeSoundByID(ctx echo.Context) error {
	userID := fmt.Sprint(ctx.Get("user_id"))
	soundID := ctx.Param("id")

	err := c.soundsService.LikeSoundByID(ctx.Request().Context(), userID, soundID)
	if err != nil {
		return fmt.Errorf("like sound by id from sounds service: %w", err)
	}

	return ctx.NoContent(http.StatusOK)
}
