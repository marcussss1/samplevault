package samplesv1

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c Controller) GetAllSamples(ctx echo.Context) error {
	sessionID := ctx.Get("session_id").(string)
	samples, err := c.samplesService.GetAllSamples(ctx.Request().Context(), sessionID)
	if err != nil {
		return fmt.Errorf("get all samples from samples service: %w", err)
	}

	return ctx.JSON(http.StatusOK, samples)
}
