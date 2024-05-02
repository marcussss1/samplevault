package samplesv1

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c Controller) GenerateSoundByImageURL(ctx echo.Context) error {
	type requestStruct struct {
		ImageURL string `json:"image_url"`
	}
	var req requestStruct
	err := ctx.Bind(&req)
	if err != nil {
		return fmt.Errorf("error while binding body: %w", err)
	}

	sound, err := c.soundsService.GenerateSoundByImageURL(ctx.Request().Context(), req.ImageURL)
	if err != nil {
		return fmt.Errorf("generate sound by image url from sounds service: %w", err)
	}

	return ctx.JSON(http.StatusOK, sound)
}
