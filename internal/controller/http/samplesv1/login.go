package samplesv1

import (
	"github.com/labstack/echo/v4"
)

func (c Controller) Login(ctx echo.Context) error {
	return nil
	//var loginUser model.LoginUser
	//err := ctx.Bind(&loginUser)
	//if err != nil {
	//	return fmt.Errorf("error while binding body: %w", err)
	//}
	//
	//user, err := c.authService.Login(ctx.Request().Context(), loginUser)
	//if err != nil {
	//	if errors.Is(err, model.ErrNotFound) {
	//		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	//	}
	//
	//	return fmt.Errorf("login from auth service: %w", err)
	//}
	//
	//ctx.SetCookie(&http.Cookie{
	//	Name:     "session_id",
	//	Value:    user.SessionID,
	//	HttpOnly: true,
	//	Path:     "/",
	//	Expires:  time.Now().Add(24 * time.Hour * 30),
	//	SameSite: http.SameSiteNoneMode,
	//	Secure:   true,
	//})
	//
	//return ctx.JSON(http.StatusOK, user)
}
