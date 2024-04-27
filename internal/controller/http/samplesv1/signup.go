package samplesv1

import (
	"fmt"
	"github.com/marcussss1/simplevault/internal/model"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func (c Controller) Signup(ctx echo.Context) error {
	var signupUser model.SignupUser
	err := ctx.Bind(&signupUser)
	if err != nil {
		return fmt.Errorf("error while binding body: %w", err)
	}

	req := ctx.Request()
	user, err := c.authService.Signup(req.Context(), signupUser)
	if err != nil {
		// todo if user exist
		return fmt.Errorf("signup from auth service: %w", err)
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

	return ctx.JSON(http.StatusCreated, user)
}
