package samplesv1

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c Controller) DownloadSound(ctx echo.Context) error {
	req := ctx.Request()
	filename := ctx.Param("filename")

	audioFile, err := c.filesService.DownloadSound(req.Context(), filename)
	if err != nil {
		return fmt.Errorf("download sounds from files service: %w", err)
	}
	defer audioFile.Close()

	return ctx.Stream(http.StatusOK, "audio/mpeg", audioFile)
}
