package templatemanager

import (
	"encoding/json"
	"net/http"
)

type HTTPRenderTemplatePresenterFactory interface {
	Create(w http.ResponseWriter) RenderTemplatePresenter
}

type RendererInteractor interface {
	RenderByJSON(template *Component, presenter RenderTemplatePresenter)
}

type RenderTemplateByJSONHandler struct {
	Interactor RendererInteractor
	Presenter  HTTPRenderTemplatePresenterFactory
}

func NewRenderTemplateByJSONHandler(
	interactor RendererInteractor,
	presenter HTTPRenderTemplatePresenterFactory,
) *RenderTemplateByJSONHandler {
	return &RenderTemplateByJSONHandler{
		Interactor: interactor,
		Presenter:  presenter,
	}
}

func (h *RenderTemplateByJSONHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req := renderTemplateByJSONRequest{}
	json.NewDecoder(r.Body).Decode(&req)

	h.Interactor.RenderByJSON(req.Template, h.Presenter.Create(w))
}

type renderTemplateByJSONRequest struct {
	Template *Component `json:"template"`
}
