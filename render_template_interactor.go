package templatemanager

import "encoding/json"

type RenderTemplatePresenter interface {
	PresentHTML(html string)
}

type Renderer interface {
	Render(Component) string
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
	presenter.PresentHTML(r.Renderer.Render(component))
}

func (r *RenderTemplateInteractor) parseJSON(templateJSON string) Component {
	component := Component{}
	json.Unmarshal([]byte(templateJSON), &component)
	return component
}
