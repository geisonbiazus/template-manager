package templatemanager

type RenderTemplatePresenter interface {
	PresentHTML(html string)
	PresentValidationErrors([]ValidationError)
}

type Renderer interface {
	Render(*Component) string
}

var invalidTemplateJSONValidationError = ValidationError{
	Field:   "template_json",
	Type:    ErrorInvalid,
	Message: "The given template JSON is invalid",
}

type RenderTemplateInteractor struct {
	Presenter RenderTemplatePresenter
	Renderer  Renderer
}

func NewRenderTemplateInteractor(renderer Renderer) *RenderTemplateInteractor {
	return &RenderTemplateInteractor{
		Renderer: renderer,
	}
}

func (r *RenderTemplateInteractor) RenderByJSON(
	template *Component, presenter RenderTemplatePresenter,
) {
	if template == nil || template.Empty() {
		presenter.PresentValidationErrors([]ValidationError{invalidTemplateJSONValidationError})
		return
	}

	presenter.PresentHTML(r.Renderer.Render(template))
}
