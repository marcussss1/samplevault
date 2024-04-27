package samplesv1

import (
	"encoding/json"
	"fmt"
	"github.com/marcussss1/simplevault/internal/model"
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

	bytes := req.FormValue("json_data")
	fmt.Println(bytes)
	var uploadSound model.UploadSound
	err = json.Unmarshal([]byte(bytes), &uploadSound)
	if err != nil {
		return fmt.Errorf("error while binding body: %w", err)
	}

	sound, err := c.filesService.UploadSound(req.Context(), audioFile, header, userID, uploadSound)
	if err != nil {
		return fmt.Errorf("upload sounds from files service: %w", err)
	}

	fmt.Println("sound: ", sound)
	return ctx.JSON(http.StatusOK, sound)
}
