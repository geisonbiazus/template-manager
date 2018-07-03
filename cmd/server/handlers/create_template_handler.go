package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/geisonbiazus/templatemanager/internal/templatemanager"
)

type ManageTemplateOutputBoundaryFactory interface {
	Create(http.ResponseWriter) templatemanager.ManageTemplateOutputBoundary
}

type CreateTemplateHandler struct {
	input         templatemanager.ManageTemplateInputBoundary
	outputFactory ManageTemplateOutputBoundaryFactory
}

func NewCreateTemplateHandler(
	input templatemanager.ManageTemplateInputBoundary,
	outputFactory ManageTemplateOutputBoundaryFactory,
) *CreateTemplateHandler {
	return &CreateTemplateHandler{
		input:         input,
		outputFactory: outputFactory,
	}
}

func (h *CreateTemplateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	request := createTemplateRequest{}
	json.NewDecoder(r.Body).Decode(&request)

	h.input.Create(request.Template, h.outputFactory.Create(w))
}

type createTemplateRequest struct {
	Template templatemanager.Template `json:"template"`
}
