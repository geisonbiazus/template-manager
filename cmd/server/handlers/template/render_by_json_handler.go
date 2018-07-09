package template

import (
	"encoding/json"
	"net/http"

	"github.com/geisonbiazus/templatemanager/internal/templatemanager"
	"github.com/geisonbiazus/templatemanager/internal/templatemanager/rendertemplate"
)

type RenderTemplateInteractor interface {
	RenderByJSON(rendertemplate.RenderByJSONRequest) rendertemplate.RenderByJSONResponse
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
	req := h.createInteractorRequest(body)
	resp := h.Interactor.RenderByJSON(req)
	h.writeSuccessResponse(w, resp)
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

func (h *RenderByJSONHandler) createInteractorRequest(
	b renderByJSONBody,
) rendertemplate.RenderByJSONRequest {
	return rendertemplate.RenderByJSONRequest{
		Template: rendertemplate.Template{
			Body: b.Template.Body,
		},
	}
}

func (h *RenderByJSONHandler) writeSuccessResponse(
	w http.ResponseWriter, resp rendertemplate.RenderByJSONResponse,
) {
	response := struct {
		HTML string `json:"html"`
	}{HTML: resp.HTML}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}