package samplesv1

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c Controller) UploadSample(ctx echo.Context) error {
	err := ctx.Request().ParseMultipartForm(5 * 1024 * 1024)
	if err != nil {
		return err
	}

	sampleFile, header, err := ctx.Request().FormFile("audio")
	if err != nil {
		return err
	}

	err = c.samplesService.UploadSample(ctx.Request().Context(), sampleFile, header)
	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusOK)
}
