package samplesv1

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (c Controller) DeleteSoundByID(ctx echo.Context) error {
	req := ctx.Request()
	id := ctx.Param("id")

	err := c.soundsService.DeleteSoundByID(req.Context(), id)
	if err != nil {
		return fmt.Errorf("delete sound by id from files service: %w", err)
	}

	return ctx.NoContent(http.StatusOK)
}
