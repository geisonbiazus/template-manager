package templatemanager

import (
	"encoding/json"
	"net/http"
)

type RenderTemplateOutputBoundaryFactory interface {
	Create(w http.ResponseWriter) RenderTemplateOutputBoundary
}

type RenderTemplateByJSONHandler struct {
	Input         RenderTemplateInputBoundary
	OutputFactory RenderTemplateOutputBoundaryFactory
}

func NewRenderTemplateByJSONHandler(
	interactor RenderTemplateInputBoundary,
	outputFactory RenderTemplateOutputBoundaryFactory,
) *RenderTemplateByJSONHandler {
	return &RenderTemplateByJSONHandler{
		Input:         interactor,
		OutputFactory: outputFactory,
	}
}

func (h *RenderTemplateByJSONHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req := renderTemplateByJSONRequest{}
	json.NewDecoder(r.Body).Decode(&req)

	h.Input.RenderByJSON(req.Template, h.OutputFactory.Create(w))
}

type renderTemplateByJSONRequest struct {
	Template *Component `json:"template"`
}
