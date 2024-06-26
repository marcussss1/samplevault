package middleware

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/marcussss1/simplevault/internal/model"
	"net/http"
	"strings"
)

func Auth(authService authService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			url := ctx.Request().URL.String()
			if url != "/api/v1/sounds/generate_by_text" &&
				url != "/api/v1/sounds/generate_by_audio_url" &&
				url != "/api/v1/sounds/generate_by_image_url" &&
				url != "/api/v1/sounds/upload" &&
				url != "/api/v1/sounds/last_generated" &&
				!strings.Contains(url, "/api/v1/sounds/download/") &&
				!strings.Contains(url, "/api/v1/sounds/like/") &&
				!strings.Contains(url, "/api/v1/liked-sounds") &&
				ctx.Request().Method != http.MethodPut &&
				ctx.Request().Method != http.MethodDelete {
				return next(ctx)
			}

			sessionID, err := ctx.Cookie("session_id")
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
			}

			user, err := authService.GetUserBySessionID(ctx.Request().Context(), sessionID.Value)
			if err != nil {
				if errors.Is(err, model.ErrNotFound) {
					return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
				}

				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}

			ctx.Set("user_id", user.ID)
			ctx.Set("session_id", sessionID.Value)

			return next(ctx)
		}
	}
}

//var users = make(map[string]string)
//
//func hash(input string) string {
//	hash := sha256.New()
//	hash.Write([]byte(input))
//	hashBytes := hash.Sum(nil)
//	hashString := hex.EncodeToString(hashBytes)
//	return hashString
//}
//
//func Auth(next echo.HandlerFunc) echo.HandlerFunc {
//	return func(ctx echo.Context) error {
//		hashUserID := hash(ctx.RealIP())
//		userID, ok := users[hashUserID]
//		if !ok {
//			users[hashUserID] = hashUserID
//		}
//
//		ctx.Set("session_id", userID)
//
//		return next(ctx)
//	}
//}
