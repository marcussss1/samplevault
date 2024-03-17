package samplesv1

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c Controller) DownloadSample(ctx echo.Context) error {
	objectReader, err := c.samplesService.DownloadSample(ctx.Request().Context())
	if err != nil {
		return err
	}
	defer objectReader.Close()

	return ctx.Stream(http.StatusOK, "audio/mpeg", objectReader)
}
