package samplesv1

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/marcussss1/simplevault/internal/model"
	"net/http"
	"time"
)

func (c Controller) Login(ctx echo.Context) error {
	var loginUserRequest model.LoginUserRequest
	err := ctx.Bind(&loginUserRequest)
	if err != nil {
		return fmt.Errorf("error while binding body: %w", err)
	}

	loginUserResponse, err := c.authService.Login(ctx.Request().Context(), loginUserRequest)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "user not found")
		}

		return fmt.Errorf("login from auth service: %w", err)
	}

	ctx.SetCookie(&http.Cookie{
		Name:     "session_id",
		Value:    loginUserResponse.SessionID,
		HttpOnly: true,
		Path:     "/",
		Expires:  time.Now().Add(24 * time.Hour * 30),
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
	})

	type responseStruct struct {
		ID       string `json:"id"`
		Username string `json:"username"`
	}
	return ctx.JSON(http.StatusOK, responseStruct{
		ID:       loginUserResponse.ID,
		Username: loginUserResponse.Username,
	})
}
