package samplesv1

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/marcussss1/simplevault/internal/model"
	"net/http"
)

func (c Controller) EditSoundByID(ctx echo.Context) error {
	var sound model.EditSound
	err := ctx.Bind(&sound)
	if err != nil {
		return fmt.Errorf("error while binding body: %w", err)
	}

	fmt.Printf("edit sound: %v", sound)

	err = c.soundsService.EditSoundByID(ctx.Request().Context(), sound)
	if err != nil {
		return fmt.Errorf("edit sound by id from sounds service: %w", err)
	}

	return ctx.NoContent(http.StatusOK)
}
