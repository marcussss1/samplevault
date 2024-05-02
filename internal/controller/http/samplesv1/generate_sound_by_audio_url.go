package samplesv1

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c Controller) GenerateSoundByAudioURL(ctx echo.Context) error {
	type requestStruct struct {
		AudioURL string `json:"audio_url"`
	}
	var req requestStruct
	err := ctx.Bind(&req)
	if err != nil {
		return fmt.Errorf("error while binding body: %w", err)
	}

	sound, err := c.soundsService.GenerateSoundByAudioURL(ctx.Request().Context(), req.AudioURL)
	if err != nil {
		return fmt.Errorf("generate sound by audio url from sounds service: %w", err)
	}

	return ctx.JSON(http.StatusOK, sound)
}
