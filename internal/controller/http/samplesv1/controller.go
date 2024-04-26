package samplesv1

import "fmt"

type Controller struct {
	soundsService    soundsService
	filesService     filesService
	playlistsService playlistsService
	authService      authService
}

func NewController(
	soundsService soundsService,
	filesService filesService,
	playlistsService playlistsService,
	authService authService,
) (*Controller, error) {
	if soundsService == nil {
		return nil, fmt.Errorf("soundsService is nil")
	}
	if filesService == nil {
		return nil, fmt.Errorf("filesService is nil")
	}
	if playlistsService == nil {
		return nil, fmt.Errorf("playlistsService is nil")
	}
	if authService == nil {
		return nil, fmt.Errorf("authService is nil")
	}

	return &Controller{
		soundsService:    soundsService,
		filesService:     filesService,
		playlistsService: playlistsService,
		authService:      authService,
	}, nil
}
