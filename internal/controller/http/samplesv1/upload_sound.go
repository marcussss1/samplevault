package samplesv1

import (
	"fmt"
	"github.com/marcussss1/simplevault/internal/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c Controller) UploadSound(ctx echo.Context) error {
	userID := fmt.Sprint(ctx.Get("user_id"))
	//userID := ctx.Request().FormValue("user_id")
	//if userID == "" {
	//}
	//
	//fmt.Println("user_id: ", userID)

	//_ = ctx.Request().ParseMultipartForm(50 * 1024 * 1024)
	////if err != nil {
	////	return fmt.Errorf("parse multipart form: %w", err)
	////}
	//
	//audioFile, header, _ := ctx.Request().FormFile("audio")
	////if header
	////if err != nil {
	////	return fmt.Errorf("form file: %w", err)
	////}
	////defer audioFile.Close()

	var req model.UploadSound
	err := ctx.Bind(&req)
	if err != nil {
		return fmt.Errorf("error while binding body: %w", err)
	}

	sound, err := c.filesService.UploadSound(ctx.Request().Context(), nil, nil, userID, req)
	if err != nil {
		return fmt.Errorf("upload sounds from files service: %w", err)
	}

	return ctx.JSON(http.StatusOK, sound)
}
