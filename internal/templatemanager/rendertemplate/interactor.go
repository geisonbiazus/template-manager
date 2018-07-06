package rendertemplate

import "github.com/geisonbiazus/templatemanager/internal/templatemanager"

const (
	StatusSuccess = "Success"
	StatusInvalid = "Invalid"
)

type Renderer interface {
	Render(*templatemanager.Component) string
}

type RenderByJSONRequest struct {
	Template *templatemanager.Component
}

type RenderByJSONResponse struct {
	Status string
	HTML   string
	Errors []templatemanager.ValidationError
}

type Interactor struct {
	Renderer Renderer
}

func NewInteractor(renderer Renderer) *Interactor {
	return &Interactor{
		Renderer: renderer,
	}
}

func (i *Interactor) RenderByJSON(r RenderByJSONRequest) RenderByJSONResponse {
	if r.Template == nil || r.Template.Empty() {
		return invalidTemplateBodyResponse
	}

	html := i.Renderer.Render(r.Template)
	return RenderByJSONResponse{Status: StatusSuccess, HTML: html}
}

var invalidTemplateBodyResponse = RenderByJSONResponse{
	Status: StatusInvalid,
	Errors: []templatemanager.ValidationError{
		templatemanager.ValidationError{
			Field:   "body",
			Type:    templatemanager.ErrorInvalid,
			Message: "The given template JSON is invalid",
		},
	},
}
