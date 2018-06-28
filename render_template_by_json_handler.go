package templatemanager

import (
	"encoding/json"
	"net/http"
)

type HTTPRenderTemplatePresenter interface {
	RenderTemplatePresenter
	With(w http.ResponseWriter) HTTPRenderTemplatePresenter
}

type RendererInteractor interface {
	RenderByJSON(template *Component, presenter RenderTemplatePresenter)
}

type RenderTemplateByJSONHandler struct {
	Interactor RendererInteractor
	Presenter  HTTPRenderTemplatePresenter
}

func NewRenderTemplateByJSONHandler(
	interactor RendererInteractor,
	presenter HTTPRenderTemplatePresenter,
) *RenderTemplateByJSONHandler {
	return &RenderTemplateByJSONHandler{
		Interactor: interactor,
		Presenter:  presenter,
	}
}

func (h *RenderTemplateByJSONHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req := renderTemplateByJSONRequest{}
	json.NewDecoder(r.Body).Decode(&req)

	h.Interactor.RenderByJSON(req.Template, h.Presenter.With(w))
}

type renderTemplateByJSONRequest struct {
	Template *Component `json:"template"`
}
