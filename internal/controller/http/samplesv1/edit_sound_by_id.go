package samplesv1

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/marcussss1/simplevault/internal/model"
	"io"
	"net/http"
)

func (c Controller) EditSoundByID(ctx echo.Context) error {
	body, err := io.ReadAll(ctx.Request().Body)
	if err != nil {
		return fmt.Errorf("read all request body: %w", err)
	}
	defer ctx.Request().Body.Close()

	var sound model.EditSound
	err = json.Unmarshal(body, &sound)
	if err != nil {
		return fmt.Errorf("decode sound 1: %w", err)
	}

	err = c.soundsService.EditSoundByID(ctx.Request().Context(), sound)
	if err != nil {
		return fmt.Errorf("edit sound by id from sounds service: %w", err)
	}

	return ctx.NoContent(http.StatusOK)
}
