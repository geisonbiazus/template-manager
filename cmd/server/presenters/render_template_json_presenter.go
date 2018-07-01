package presenters

import (
	"encoding/json"
	"net/http"

	"github.com/geisonbiazus/templatemanager"
)

type RenderTemplateJSONPresenter struct {
	Writer http.ResponseWriter
}

func (p *RenderTemplateJSONPresenter) PresentHTML(html string) {
	result := struct {
		HTML string `json:"html"`
	}{
		HTML: html,
	}

	p.renderJSON(http.StatusOK, result)
}

func (p *RenderTemplateJSONPresenter) PresentValidationErrors(errors []templatemanager.ValidationError) {
	result := struct {
		Errors []templatemanager.ValidationError `json:"errors"`
	}{
		Errors: errors,
	}
	p.renderJSON(http.StatusUnprocessableEntity, result)
}

func (p *RenderTemplateJSONPresenter) renderJSON(status int, body interface{}) {
	p.Writer.Header().Set("Content-Type", "application/json")
	p.Writer.WriteHeader(status)
	json.NewEncoder(p.Writer).Encode(body)
}

type RenderTemplateJSONPresenterFactory struct{}

func NewRenderTemplateJSONPresenterFactory() *RenderTemplateJSONPresenterFactory {
	return &RenderTemplateJSONPresenterFactory{}
}

func (f *RenderTemplateJSONPresenterFactory) Create(
	w http.ResponseWriter,
) templatemanager.RenderTemplateOutputBoundary {
	return &RenderTemplateJSONPresenter{Writer: w}
}
