package middleware

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/labstack/echo/v4"
)

// TODO разумеется временно
// в тарантуле добавить хранилку сессий
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
		//cookie, err := ctx.Cookie("session_id")
		//if err != nil {
		//	sessionID = hash(ctx.RealIP())
		//	ctx.SetCookie(&http.Cookie{
		//		Name:     "session_id",
		//		Value:    sessionID,
		//		Path:     "/",
		//		Expires:  time.Now().Add(24 * time.Hour * 30),
		//		Secure:   true,
		//		HttpOnly: true,
		//		SameSite: http.SameSiteNoneMode,
		//	})
		//} else {
		//	sessionID = cookie.Value
		//}

		ctx.Set("session_id", userID)

		return next(ctx)
	}
}
