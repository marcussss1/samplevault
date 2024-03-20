package samplesv1

import "fmt"

type Controller struct {
	soundsService soundsService
	filesService  filesService
}

func NewController(
	soundsService soundsService,
	filesService filesService,
) (*Controller, error) {
	if soundsService == nil {
		return nil, fmt.Errorf("soundsService is nil")
	}
	if filesService == nil {
		return nil, fmt.Errorf("filesService is nil")
	}

	return &Controller{
		soundsService: soundsService,
		filesService:  filesService,
	}, nil
}
