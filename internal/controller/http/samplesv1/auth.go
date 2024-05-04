package samplesv1

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/marcussss1/simplevault/internal/model"
	"net/http"
)

func (c Controller) Auth(ctx echo.Context) error {
	sessionID, err := ctx.Cookie("session_id")
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "cookie is empty")
	}

	user, err := c.authService.GetUserBySessionID(ctx.Request().Context(), sessionID.Value)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	fmt.Println(user)

	return ctx.JSON(http.StatusOK, user)
}
