package samplesv1

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c Controller) GetLastGeneratedUserSounds(ctx echo.Context) error {
	userID := fmt.Sprint(ctx.Get("user_id"))

	fmt.Println("user_id: ", userID)

	sounds, err := c.soundsService.GetLastGeneratedUserSounds(ctx.Request().Context(), userID)
	if err != nil {
		return fmt.Errorf("get last generated user sounds from sounds service: %w", err)
	}

	fmt.Println("sounds: ", sounds)

	return ctx.JSON(http.StatusOK, sounds)
}
