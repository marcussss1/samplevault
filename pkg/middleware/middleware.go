package middleware

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/labstack/echo/v4"
)

//func Auth(authService authService) echo.MiddlewareFunc {
//	return func(next echo.HandlerFunc) echo.HandlerFunc {
//		return func(ctx echo.Context) error {
//			url := ctx.Request().URL.String()
//			if url != "/api/v1/sounds/upload" {
//				return next(ctx)
//			}
//
//			sessionID, err := ctx.Cookie("session_id")
//			if err != nil {
//				return echo.NewHTTPError(http.StatusUnauthorized, "cookie is empty")
//			}
//
//			user, err := authService.GetUserBySessionID(ctx.Request().Context(), sessionID.String())
//			if err != nil {
//				if errors.Is(err, model.ErrNotFound) {
//					return echo.NewHTTPError(http.StatusUnauthorized, "user not found")
//				}
//
//				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
//			}
//
//			ctx.Set("user", user)
//			return next(ctx)
//		}
//	}
//}

var users = make(map[string]string)

func hash(input string) string {
	hash := sha256.New()
	hash.Write([]byte(input))
	hashBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	return hashString
}

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		hashUserID := hash(ctx.RealIP())
		userID, ok := users[hashUserID]
		if !ok {
			users[hashUserID] = hashUserID
		}

		ctx.Set("session_id", userID)

		return next(ctx)
	}
}
