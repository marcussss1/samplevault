package samplesv1

import "fmt"

type Controller struct {
	samplesService samplesService
}

func NewController(
	samplesService samplesService,
) (*Controller, error) {
	if samplesService == nil {
		return nil, fmt.Errorf("samplesService is nil")
	}

	return &Controller{
		samplesService: samplesService,
	}, nil
}
