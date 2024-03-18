package samplesv1

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c Controller) UploadSample(ctx echo.Context) error {
	req := ctx.Request()
	// TODO Временно, когда появится авторизация нужно поменять логику
	userID := ctx.Get("session_id").(string)

	err := req.ParseMultipartForm(5 * 1024 * 1024)
	if err != nil {
		return fmt.Errorf("parse multipart form: %w", err)
	}

	sampleFile, header, err := req.FormFile("audio")
	if err != nil {
		return fmt.Errorf("form file: %w", err)
	}

	sample, err := c.filesService.UploadSample(req.Context(), sampleFile, header, userID)
	if err != nil {
		return fmt.Errorf("upload sample from samples service: %w", err)
	}

	return ctx.JSON(http.StatusOK, sample)
}
