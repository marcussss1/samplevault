package samplesv1

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (c Controller) GenerateSample(ctx echo.Context) error {
	objectReader, err := c.samplesService.GenerateSample(ctx.Request().Context())
	if err != nil {
		return err
	}
	defer objectReader.Close()

	return ctx.Stream(http.StatusOK, "audio/mpeg", objectReader)
}
