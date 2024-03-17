package samplesv1

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (c Controller) GetUrl(ctx echo.Context) error {
	url, err := c.samplesService.GetUrl()
	if err != nil {
		return fmt.Errorf("get url from samples service: %w", err)
	}

	return ctx.JSON(http.StatusOK, url)
}
