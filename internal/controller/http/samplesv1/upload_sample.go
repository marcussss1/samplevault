package samplesv1

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (c Controller) UploadSample(ctx echo.Context) error {
	maxSize := int64(64 << 20)
	err := ctx.Request().ParseMultipartForm(maxSize)
	if err != nil {
		return err
	}

	file, header, err := ctx.Request().FormFile("audio")
	if err != nil {
		return err
	}

	err = c.samplesService.UploadSample(ctx.Request().Context(), file, header)
	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusOK)
}
