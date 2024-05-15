package samplesv1

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c Controller) GenerateSoundByAudioURL(ctx echo.Context) error {
	userID := fmt.Sprint(ctx.Get("user_id"))

	fmt.Println("user_id: ", userID)

	//type requestStruct struct {
	//	AudioURL string `json:"audio_url"`
	//}
	//var req requestStruct
	//err := ctx.Bind(&req)
	//if err != nil {
	//	return fmt.Errorf("error while binding body: %w", err)
	//}

	//audioFile, header, err := ctx.Request().FormFile("audio")
	//if err != nil {
	//	return fmt.Errorf("form file: %w", err)
	//}
	//defer audioFile.Close()
	//
	//sound, err := c.soundsService.GenerateSoundByAudioURL(ctx.Request().Context(), audioFile, header, userID)
	//if err != nil {
	//	return fmt.Errorf("generate sound by audio url from sounds service: %w", err)
	//}

	sound, err := c.soundsService.GenerateSoundByText(ctx.Request().Context(), "generate a sad sequence of notes played on the piano, the mood is dark and sad", userID)
	if err != nil {
		return fmt.Errorf("generate sound by text from sounds service: %w", err)
	}

	return ctx.JSON(http.StatusCreated, sound)
}
