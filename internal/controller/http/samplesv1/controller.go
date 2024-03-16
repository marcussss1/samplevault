package samplesv1

type Controller struct {
	samplesService samplesService
}

func NewController(
	samplesService samplesService,
) *Controller {
	return &Controller{
		samplesService: samplesService,
	}
}
