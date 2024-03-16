package samplesv1

import (
	"github.com/labstack/echo/v4"
	"io"
	"os"
)

func (c Controller) DownloadSample(ctx echo.Context) error {
	objectReader, err := c.samplesService.DownloadSample(ctx.Request().Context())
	if err != nil {
		return err
	}
	defer objectReader.Close()

	file, err := os.Create("sample.mp3")
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, objectReader)
	if err != nil {
		return err
	}

	return ctx.File("sample.mp3")
}
