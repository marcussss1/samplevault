package samplesv1

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (c Controller) GetRandomPlaylists(ctx echo.Context) error {
	// TODO Временно, когда появится авторизация нужно поменять логику
	req := ctx.Request()
	userID := ctx.Get("session_id").(string)
	fmt.Println("userID: ", userID)
	playlists, err := c.playlistsService.GetRandomPlaylists(req.Context())
	if err != nil {
		return fmt.Errorf("get random playlists from playlists service: %w", err)
	}
	fmt.Println("playlists: ", playlists)
	return ctx.JSON(http.StatusOK, playlists)
}
