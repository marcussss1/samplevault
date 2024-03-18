package samplesv1

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c Controller) GenerateSample(ctx echo.Context) error {
	sampleFile, err := c.filesService.GenerateSample(ctx.Request().Context())
	if err != nil {
		return fmt.Errorf("generate sample from samples service: %w", err)
	}
	defer sampleFile.Close()

	return ctx.Stream(http.StatusOK, "audio/mpeg", sampleFile)
}
