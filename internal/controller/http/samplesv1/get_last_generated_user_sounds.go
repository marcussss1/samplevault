package samplesv1

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c Controller) GetLastGeneratedUserSounds(ctx echo.Context) error {
	// TODO Временно, когда появится авторизация нужно поменять логику
	req := ctx.Request()
	userID := ctx.Get("session_id").(string)
	fmt.Println("userID: ", userID)
	sounds, err := c.soundsService.GetLastGeneratedUserSounds(req.Context(), userID)
	if err != nil {
		return fmt.Errorf("get last generated user sounds from sounds service: %w", err)
	}
	fmt.Println("sounds: ", sounds)
	return ctx.JSON(http.StatusOK, sounds)
}
