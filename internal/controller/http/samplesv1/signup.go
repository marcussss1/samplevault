package samplesv1

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/marcussss1/simplevault/internal/model"
	"net/http"
	"time"
)

func (c Controller) Signup(ctx echo.Context) error {
	var signupUserRequest model.SignupUserRequest
	err := ctx.Bind(&signupUserRequest)
	if err != nil {
		return fmt.Errorf("error while binding body: %w", err)
	}

	signupUserResponse, err := c.authService.Signup(ctx.Request().Context(), signupUserRequest)
	if err != nil {
		// todo if user exist
		return fmt.Errorf("signup from auth service: %w", err)
	}

	ctx.SetCookie(&http.Cookie{
		Name:     "session_id",
		Value:    signupUserResponse.SessionID,
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
	return ctx.JSON(http.StatusCreated, responseStruct{
		ID:       signupUserResponse.ID,
		Username: signupUserResponse.Username,
	})
}
