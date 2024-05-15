package samplesv1

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c Controller) GenerateSoundByImageURL(ctx echo.Context) error {
	userID := fmt.Sprint(ctx.Get("user_id"))

	fmt.Println("user_id: ", userID)

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

	sound, err := c.soundsService.GenerateSoundByText(ctx.Request().Context(), "nice ocean, great blue beer amazing nice great nice girls mountains song sign", userID)
	if err != nil {
		return fmt.Errorf("generate sound by text from sounds service: %w", err)
	}

	return ctx.JSON(http.StatusCreated, sound)
}
