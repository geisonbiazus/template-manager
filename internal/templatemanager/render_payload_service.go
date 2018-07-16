package templatemanager

type ContentPresenter interface {
	PresentContent(content string)
}

type Renderer interface {
	Render(c *Component) string
}

type RenderPayloadService struct {
	Component *Component
	Renderer  Renderer
	Presenter ContentPresenter
}

func NewRenderPayloadService(
	component *Component, renderer Renderer, presenter ContentPresenter,
) *RenderPayloadService {

	return &RenderPayloadService{
		Component: component,
		Renderer:  renderer,
		Presenter: presenter,
	}
}

func (r *RenderPayloadService) Execute() {
	r.Presenter.PresentContent(r.Renderer.Render(r.Component))
}
