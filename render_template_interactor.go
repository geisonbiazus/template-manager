package templatemanager

import (
	"encoding/json"
)

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
	templateJSON string, presenter RenderTemplatePresenter,
) {
	component := r.parseJSON(templateJSON)

	if component == nil {
		presenter.PresentValidationErrors([]ValidationError{invalidTemplateJSONValidationError})
		return
	}

	presenter.PresentHTML(r.Renderer.Render(component))
}

func (r *RenderTemplateInteractor) parseJSON(templateJSON string) *Component {
	component := &Component{}
	err := json.Unmarshal([]byte(templateJSON), component)
	if err != nil || component.Empty() {
		return nil
	}
	return component
}
