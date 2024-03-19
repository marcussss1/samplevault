package samplesv1

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c Controller) GetAllSamples(ctx echo.Context) error {
	// TODO Временно, когда появится авторизация нужно поменять логику
	req := ctx.Request()
	userID := ctx.Get("session_id").(string)
	fmt.Println("userID: ", userID)
	samples, err := c.samplesService.GetAllSamples(req.Context(), userID)
	if err != nil {
		return fmt.Errorf("get all samples from samples service: %w", err)
	}

	return ctx.JSON(http.StatusOK, samples)
}
