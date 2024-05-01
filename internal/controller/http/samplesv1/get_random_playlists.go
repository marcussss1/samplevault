package samplesv1

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (c Controller) GetRandomPlaylists(ctx echo.Context) error {
	playlists, err := c.playlistsService.GetRandomPlaylists(ctx.Request().Context())
	if err != nil {
		return fmt.Errorf("get random playlists from playlists service: %w", err)
	}

	fmt.Println("playlists: ", playlists)

	return ctx.JSON(http.StatusOK, playlists)
}
