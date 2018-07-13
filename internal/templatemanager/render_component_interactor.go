package templatemanager

type ContentPresenter interface {
}

type RenderPayloadInput struct {
	Component *Component
}

type RenderPayloadInteractor struct {
}

func NewRenderPayloadInteractor(input RenderPayloadInput, presenter ContentPresenter) *RenderPayloadInteractor {
	return &RenderPayloadInteractor{}
}

func (r *RenderPayloadInteractor) Execute() {

}
