package samplesv1

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c Controller) UploadSound(ctx echo.Context) error {
	req := ctx.Request()
	// TODO Временно, когда появится авторизация нужно поменять логику
	userID := ctx.Get("session_id").(string)
	fmt.Println("userID: ", userID)
	err := req.ParseMultipartForm(5 * 1024 * 1024)
	if err != nil {
		return fmt.Errorf("parse multipart form: %w", err)
	}

	audioFile, header, err := req.FormFile("audio")
	if err != nil {
		return fmt.Errorf("form file: %w", err)
	}
	defer audioFile.Close()

	sound, err := c.filesService.UploadSound(req.Context(), audioFile, header, userID)
	if err != nil {
		return fmt.Errorf("upload sounds from files service: %w", err)
	}

	fmt.Println("sound: ", sound)
	return ctx.JSON(http.StatusOK, sound)
}
