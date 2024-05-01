package samplesv1

import (
	"fmt"
	"github.com/marcussss1/simplevault/internal/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c Controller) GetLastGeneratedUserSounds(ctx echo.Context) error {
	// TODO Временно, когда появится авторизация нужно поменять логику
	user := ctx.Get("user").(model.User)

	fmt.Println("user: ", user)

	sounds, err := c.soundsService.GetLastGeneratedUserSounds(ctx.Request().Context(), user.ID)
	if err != nil {
		return fmt.Errorf("get last generated user sounds from sounds service: %w", err)
	}

	fmt.Println("sounds: ", sounds)

	return ctx.JSON(http.StatusOK, sounds)
}
