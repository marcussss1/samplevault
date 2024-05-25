package samplesv1

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c Controller) GenerateSoundByImageURL(ctx echo.Context) error {
	userID := fmt.Sprint(ctx.Get("user_id"))
	sessionID := fmt.Sprint(ctx.Get("session_id"))

	//imageFile, header, err := ctx.Request().FormFile("audio")
	//if err != nil {
	//	return fmt.Errorf("form file: %w", err)
	//}
	//defer imageFile.Close()
	//
	//sound, err := c.soundsService.GenerateSoundByImageURL(ctx.Request().Context(), imageFile, header, userID)
	//if err != nil {
	//	return fmt.Errorf("generate sound by image url from sounds service: %w", err)
	//}

	sound, err := c.soundsService.GenerateSoundByText(ctx.Request().Context(), "sounds of the sea with smooth waves as a rhythmic and the sounds of the surf as a melodic accent. Soft guitar chords with lightness of the sea wind. Endless harmony with nature and tranquility", "8", userID, sessionID)
	if err != nil {
		return fmt.Errorf("generate sound by text from sounds service: %w", err)
	}

	return ctx.JSON(http.StatusCreated, sound)
}
