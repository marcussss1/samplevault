package samplesv1

import (
	"fmt"
	"github.com/marcussss1/simplevault/internal/model"
	"net/http"

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
		return fmt.Errorf("signup from auth service: %w", err)
	}

	return ctx.JSON(http.StatusCreated, user)
}
