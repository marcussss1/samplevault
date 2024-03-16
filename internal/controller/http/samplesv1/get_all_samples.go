package samplesv1

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (c Controller) GetAllSamples(ctx echo.Context) error {
	samples, err := c.samplesService.GetAllSamples(ctx.Request().Context(), "a2802d62-b006-4949-8fa0-07328bd26719")
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, samples)
}
