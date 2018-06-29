package templatemanager

type Renderer interface {
	Render(*Component) string
}

var invalidTemplateJSONValidationError = ValidationError{
	Field:   "template_json",
	Type:    ErrorInvalid,
	Message: "The given template JSON is invalid",
}

type RenderTemplateInteractor struct {
	Presenter RenderTemplateOutputBoundary
	Renderer  Renderer
}

func NewRenderTemplateInteractor(renderer Renderer) *RenderTemplateInteractor {
	return &RenderTemplateInteractor{
		Renderer: renderer,
	}
}

func (r *RenderTemplateInteractor) RenderByJSON(
	template *Component, presenter RenderTemplateOutputBoundary,
) {
	if template == nil || template.Empty() {
		presenter.PresentValidationErrors([]ValidationError{invalidTemplateJSONValidationError})
		return
	}

	presenter.PresentHTML(r.Renderer.Render(template))
}
