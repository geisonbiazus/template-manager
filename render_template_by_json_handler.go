package templatemanager

import (
	"encoding/json"
	"net/http"
)

type RenderTemplateOutputBoundaryFactory interface {
	Create(w http.ResponseWriter) RenderTemplateOutputBoundary
}

type RenderTemplateByJSONHandler struct {
	Interactor       RenderTemplateInputBoundary
	PresenterFactory RenderTemplateOutputBoundaryFactory
}

func NewRenderTemplateByJSONHandler(
	interactor RenderTemplateInputBoundary,
	presenterFactory RenderTemplateOutputBoundaryFactory,
) *RenderTemplateByJSONHandler {
	return &RenderTemplateByJSONHandler{
		Interactor:       interactor,
		PresenterFactory: presenterFactory,
	}
}

func (h *RenderTemplateByJSONHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req := renderTemplateByJSONRequest{}
	json.NewDecoder(r.Body).Decode(&req)

	h.Interactor.RenderByJSON(req.Template, h.PresenterFactory.Create(w))
}

type renderTemplateByJSONRequest struct {
	Template *Component `json:"template"`
}
