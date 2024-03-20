package samplesv1

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c Controller) GenerateSound(ctx echo.Context) error {
	audioFile, err := c.filesService.GenerateSound(ctx.Request().Context())
	if err != nil {
		return fmt.Errorf("generate sound from sounds service: %w", err)
	}
	defer audioFile.Close()

	return ctx.Stream(http.StatusOK, "audio/mpeg", audioFile)
}
