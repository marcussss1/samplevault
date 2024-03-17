package samplesv1

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c Controller) GenerateSample(ctx echo.Context) error {
	objectReader, err := c.samplesService.GenerateSample(ctx.Request().Context())
	if err != nil {
		return err
	}
	defer objectReader.Close()

	return ctx.Stream(http.StatusOK, "audio/mpeg", objectReader)
}
