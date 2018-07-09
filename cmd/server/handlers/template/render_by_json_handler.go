package template

import (
	"encoding/json"
	"net/http"

	"github.com/geisonbiazus/templatemanager/internal/templatemanager"
	"github.com/geisonbiazus/templatemanager/internal/templatemanager/rendertemplate"
)

type RenderTemplateInteractor interface {
	RenderByJSON(rendertemplate.RenderByJSONInput) rendertemplate.RenderByJSONOutput
}

type RenderByJSONHandler struct {
	Interactor RenderTemplateInteractor
}

func NewRenderByJSONHandler(interactor RenderTemplateInteractor) *RenderByJSONHandler {
	return &RenderByJSONHandler{
		Interactor: interactor,
	}
}

func (h *RenderByJSONHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body := h.decodeRequestBody(r)
	input := h.createInteractorInput(body)
	output := h.Interactor.RenderByJSON(input)
	h.writeResponse(w, output)
}

type renderByJSONBody struct {
	Template struct {
		Body *templatemanager.Component `json:"body"`
	} `json:"template"`
}

func (h *RenderByJSONHandler) decodeRequestBody(r *http.Request) renderByJSONBody {
	body := renderByJSONBody{}
	json.NewDecoder(r.Body).Decode(&body)
	return body
}

func (h *RenderByJSONHandler) createInteractorInput(
	b renderByJSONBody,
) rendertemplate.RenderByJSONInput {
	return rendertemplate.RenderByJSONInput{
		Template: b.Template.Body,
	}
}

func (h *RenderByJSONHandler) writeResponse(
	w http.ResponseWriter, output rendertemplate.RenderByJSONOutput,
) {
	if output.Status == rendertemplate.StatusSuccess {
		h.writeSuccessResponse(w, output)
	} else {
		h.writeInvalidResponse(w, output)
	}
}

func (h *RenderByJSONHandler) writeSuccessResponse(
	w http.ResponseWriter, output rendertemplate.RenderByJSONOutput,
) {
	body := struct {
		HTML string `json:"html"`
	}{HTML: output.HTML}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(body)
}

func (h *RenderByJSONHandler) writeInvalidResponse(
	w http.ResponseWriter, output rendertemplate.RenderByJSONOutput,
) {
	body := struct {
		Errors []templatemanager.ValidationError `json:"errors"`
	}{Errors: output.Errors}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnprocessableEntity)
	json.NewEncoder(w).Encode(body)
}
