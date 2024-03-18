package samplesv1

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c Controller) DownloadSample(ctx echo.Context) error {
	req := ctx.Request()
	filename := ctx.Param("filename")

	sampleFile, err := c.filesService.DownloadSample(req.Context(), filename)
	if err != nil {
		return fmt.Errorf("download sample from samples service: %w", err)
	}
	defer sampleFile.Close()

	return ctx.Stream(http.StatusOK, "audio/mpeg", sampleFile)
}
