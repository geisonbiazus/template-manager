package templatemanager

import "encoding/json"

type RenderTemplatePresenter interface {
	PresentHTML(html string)
}

type RenderTemplateInteractor struct {
	Presenter RenderTemplatePresenter
}

func NewRenderTemplateInteractor() *RenderTemplateInteractor {
	return &RenderTemplateInteractor{}
}

func (r *RenderTemplateInteractor) RenderByJSON(
	templateJSON string, presenter RenderTemplatePresenter,
) {
	component := r.parseJSON(templateJSON)
	presenter.PresentHTML(component.Render())
}

func (r *RenderTemplateInteractor) parseJSON(templateJSON string) Component {
	component := Component{}
	json.Unmarshal([]byte(templateJSON), &component)
	return component
}
