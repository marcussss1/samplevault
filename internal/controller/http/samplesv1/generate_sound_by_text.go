package samplesv1

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c Controller) GenerateSoundByText(ctx echo.Context) error {
	userID := fmt.Sprint(ctx.Get("user_id"))
	sessionID := fmt.Sprint(ctx.Get("session_id"))

	type requestStruct struct {
		Text     string `json:"text"`
		Duration string `json:"duration"`
	}
	var req requestStruct
	err := ctx.Bind(&req)
	if err != nil {
		return fmt.Errorf("error while binding body: %w", err)
	}

	sound, err := c.soundsService.GenerateSoundByText(ctx.Request().Context(), req.Text, req.Duration, userID, sessionID)
	if err != nil {
		return fmt.Errorf("generate sound by text from sounds service: %w", err)
	}

	return ctx.JSON(http.StatusCreated, sound)
}
