package samplesv1

import (
	"fmt"
	"github.com/marcussss1/simplevault/internal/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c Controller) UploadSound(ctx echo.Context) error {
	userID := ctx.Request().FormValue("user_id")
	if userID == "" {
		userID = fmt.Sprint(ctx.Get("user_id"))
	}

	fmt.Println("user_id: ", userID)

	err := ctx.Request().ParseMultipartForm(50 * 1024 * 1024)
	if err != nil {
		return fmt.Errorf("parse multipart form: %w", err)
	}

	audioFile, header, err := ctx.Request().FormFile("audio")
	if err != nil {
		return fmt.Errorf("form file: %w", err)
	}
	defer audioFile.Close()

	sound, err := c.filesService.UploadSound(ctx.Request().Context(), audioFile, header, userID, model.UploadSound{
		AuthorID:          userID,
		AudioURL:          ctx.Request().FormValue("audio_url"),
		Title:             ctx.Request().FormValue("title"),
		MusicalInstrument: ctx.Request().FormValue("musical_instrument"),
		Genre:             ctx.Request().FormValue("genre"),
		Mood:              ctx.Request().FormValue("mood"),
		Tonality:          ctx.Request().FormValue("tonality"),
		Tempo:             ctx.Request().FormValue("tempo"),
		IsGenerated:       ctx.Request().FormValue("is_generated"),
	})
	if err != nil {
		return fmt.Errorf("upload sounds from files service: %w", err)
	}

	fmt.Println("sound: ", sound)

	return ctx.JSON(http.StatusOK, sound)
}
