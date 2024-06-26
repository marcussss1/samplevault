package router

import (
	"github.com/labstack/echo/v4"
	samplescontrollerv1 "github.com/marcussss1/simplevault/internal/controller/http/samplesv1"
	"net/http"
)

func Route(e *echo.Echo, ctrl *samplescontrollerv1.Controller) *echo.Echo {
	e.GET("/api/v1/sounds", ctrl.GetAllSounds)
	e.POST("/api/v1/sounds/generate_by_text", ctrl.GenerateSoundByText)
	e.POST("/api/v1/sounds/generate_by_audio_url", ctrl.GenerateSoundByAudioURL)
	e.POST("/api/v1/sounds/generate_by_image_url", ctrl.GenerateSoundByImageURL)
	e.GET("/api/v1/sounds/download/:filename", ctrl.DownloadSound)
	e.POST("/api/v1/sounds/upload", ctrl.UploadSound)
	e.GET("/api/v1/sounds/random", ctrl.GetRandomSounds)
	e.GET("/api/v1/sounds/last_generated", ctrl.GetLastGeneratedUserSounds)
	e.DELETE("/api/v1/sounds/:id", ctrl.DeleteSoundByID)
	e.PUT("/api/v1/sounds", ctrl.EditSoundByID)
	e.POST("/api/v1/sounds/like/:id", ctrl.LikeSoundByID)
	e.POST("/api/v1/sounds/dislike/:id", ctrl.DislikeSoundByID)
	e.GET("/api/v1/liked-sounds", ctrl.GetLikedSounds)

	e.GET("/api/v1/playlists/random", ctrl.GetRandomPlaylists)
	e.GET("/api/v1/playlists/:playlist_id", ctrl.GetPlaylist)

	e.GET("api/v1/auth", ctrl.Auth)
	e.GET("api/v1/auth/logout", ctrl.Logout)
	e.POST("api/v1/auth/signup", ctrl.Signup)
	e.POST("api/v1/auth/login", ctrl.Login)

	e.GET("/api/v1/health", func(ctx echo.Context) error {
		return ctx.NoContent(http.StatusOK)
	})
	e.GET("/api/v1/read", func(ctx echo.Context) error {
		return ctx.NoContent(http.StatusOK)
	})

	return e
}
