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
	Renderer Renderer
}

func NewRenderTemplateInteractor(renderer Renderer) *RenderTemplateInteractor {
	return &RenderTemplateInteractor{
		Renderer: renderer,
	}
}

func (r *RenderTemplateInteractor) RenderByJSON(
	template *Component, output RenderTemplateOutputBoundary,
) {
	if template == nil || template.Empty() {
		output.PresentValidationErrors([]ValidationError{invalidTemplateJSONValidationError})
		return
	}

	output.PresentHTML(r.Renderer.Render(template))
}
