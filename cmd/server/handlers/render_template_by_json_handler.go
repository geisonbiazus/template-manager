package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/geisonbiazus/templatemanager/internal/templatemanager"
)

type RenderTemplateOutputBoundaryFactory interface {
	Create(w http.ResponseWriter) templatemanager.RenderTemplateOutputBoundary
}

type RenderTemplateByJSONHandler struct {
	Input         templatemanager.RenderTemplateInputBoundary
	OutputFactory RenderTemplateOutputBoundaryFactory
}

func NewRenderTemplateByJSONHandler(
	input templatemanager.RenderTemplateInputBoundary,
	outputFactory RenderTemplateOutputBoundaryFactory,
) *RenderTemplateByJSONHandler {
	return &RenderTemplateByJSONHandler{
		Input:         input,
		OutputFactory: outputFactory,
	}
}

func (h *RenderTemplateByJSONHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req := renderTemplateByJSONRequest{}
	json.NewDecoder(r.Body).Decode(&req)

	h.Input.RenderByJSON(req.Template, h.OutputFactory.Create(w))
}

type renderTemplateByJSONRequest struct {
	Template *templatemanager.Component `json:"template"`
}
