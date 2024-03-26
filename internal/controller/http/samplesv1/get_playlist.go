package samplesv1

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c Controller) GetPlaylist(ctx echo.Context) error {
	req := ctx.Request()
	id := ctx.Param("playlist_id")

	playlist, err := c.playlistsService.GetPlaylist(req.Context(), id)
	if err != nil {
		return fmt.Errorf("get playlist from playlists service: %w", err)
	}

	return ctx.JSON(http.StatusOK, playlist)
}
