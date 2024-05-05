package samplesv1

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/marcussss1/simplevault/internal/model"
	"net/http"
	"time"
)

func (c Controller) Logout(ctx echo.Context) error {
	sessionID, err := ctx.Cookie("session_id")
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "cookie is empty")
	}

	err = c.authService.DeleteSessionByID(ctx.Request().Context(), sessionID.Value)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}

		return fmt.Errorf("delete session by id from auth service: %w", err)
	}

	ctx.SetCookie(&http.Cookie{
		Name:     "session_id",
		Value:    "",
		HttpOnly: true,
		Expires:  time.Now().AddDate(0, 0, -1),
		Path:     "/",
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
	})

	return ctx.NoContent(http.StatusOK)
}
