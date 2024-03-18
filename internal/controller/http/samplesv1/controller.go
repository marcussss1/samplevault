package samplesv1

import "fmt"

type Controller struct {
	samplesService samplesService
	filesService   filesService
}

func NewController(
	samplesService samplesService,
	filesService filesService,
) (*Controller, error) {
	if samplesService == nil {
		return nil, fmt.Errorf("samplesService is nil")
	}
	if filesService == nil {
		return nil, fmt.Errorf("filesService is nil")
	}

	return &Controller{
		samplesService: samplesService,
		filesService:   filesService,
	}, nil
}
