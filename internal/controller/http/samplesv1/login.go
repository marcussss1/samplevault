package samplesv1

import (
	"errors"
	"fmt"
	"github.com/marcussss1/simplevault/internal/model"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func (c Controller) Login(ctx echo.Context) error {
	var loginUser model.LoginUser
	err := ctx.Bind(&loginUser)
	if err != nil {
		return fmt.Errorf("error while binding body: %w", err)
	}

	req := ctx.Request()
	user, err := c.authService.Login(req.Context(), loginUser)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return ctx.NoContent(http.StatusNotFound)
		}

		return fmt.Errorf("login from auth service: %w", err)
	}

	ctx.SetCookie(&http.Cookie{
		Name:     "session_id",
		Value:    user.SessionID,
		HttpOnly: true,
		Path:     "/",
		Expires:  time.Now().Add(24 * time.Hour * 30),
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
	})

	return ctx.JSON(http.StatusOK, user)
}
