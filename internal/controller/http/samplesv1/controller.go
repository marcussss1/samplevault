package samplesv1

import "fmt"

type Controller struct {
	soundsService    soundsService
	filesService     filesService
	playlistsService playlistsService
}

func NewController(
	soundsService soundsService,
	filesService filesService,
	playlistsService playlistsService,
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

	return &Controller{
		soundsService:    soundsService,
		filesService:     filesService,
		playlistsService: playlistsService,
	}, nil
}
