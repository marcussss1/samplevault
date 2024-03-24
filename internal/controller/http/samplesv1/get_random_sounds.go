package samplesv1

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (c Controller) GetRandomSounds(ctx echo.Context) error {
	// TODO Временно, когда появится авторизация нужно поменять логику
	req := ctx.Request()
	userID := ctx.Get("session_id").(string)
	fmt.Println("userID: ", userID)
	sounds, err := c.soundsService.GetRandomSounds(req.Context())
	if err != nil {
		return fmt.Errorf("get random sounds from sounds service: %w", err)
	}
	fmt.Println("sounds: ", sounds)
	return ctx.JSON(http.StatusOK, sounds)
}
