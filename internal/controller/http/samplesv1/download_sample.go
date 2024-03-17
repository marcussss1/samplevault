package samplesv1

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c Controller) DownloadSample(ctx echo.Context) error {
	sampleFile, err := c.samplesService.DownloadSample(ctx.Request().Context())
	if err != nil {
		return err
	}
	defer sampleFile.Close()

	return ctx.Stream(http.StatusOK, "audio/mpeg", sampleFile)
}
