package samplesv1

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c Controller) GenerateSoundByText(ctx echo.Context) error {
	type requestStruct struct {
		Text string `json:"text"`
	}
	var req requestStruct
	err := ctx.Bind(&req)
	if err != nil {
		return fmt.Errorf("error while binding body: %w", err)
	}

	sound, err := c.soundsService.GenerateSoundByText(ctx.Request().Context(), req.Text)
	if err != nil {
		return fmt.Errorf("generate sound by text from sounds service: %w", err)
	}

	return ctx.JSON(http.StatusCreated, sound)
}
