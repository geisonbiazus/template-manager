package templatemanager

type ContentPresenter interface {
	PresentContent(content string)
	PresentValidationErrors(errs []ValidationError)
}

type Renderer interface {
	Render(c *Component) string
}

type RenderPayloadService struct {
	Payload   *Component
	Renderer  Renderer
	Presenter ContentPresenter
}

func NewRenderPayloadService(
	payload *Component, renderer Renderer, presenter ContentPresenter,
) *RenderPayloadService {

	return &RenderPayloadService{
		Payload:   payload,
		Renderer:  renderer,
		Presenter: presenter,
	}
}

func (r *RenderPayloadService) Execute() {
	if r.Payload == nil || r.Payload.Empty() {
		r.Presenter.PresentValidationErrors(invalidPayloadErrors)
		return
	}

	r.Presenter.PresentContent(r.Renderer.Render(r.Payload))
}

var invalidPayloadErrors = []ValidationError{
	ValidationError{Field: "payload", Type: "INVALID", Message: "Invalid payload"},
}
