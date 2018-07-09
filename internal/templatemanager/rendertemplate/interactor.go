package rendertemplate

import (
	"github.com/geisonbiazus/templatemanager/internal/templatemanager"
)

const (
	StatusSuccess = "Success"
	StatusInvalid = "Invalid"
)

type Renderer interface {
	Render(*templatemanager.Component) string
}

type RenderByJSONInput struct {
	Template *templatemanager.Component
}

type RenderByJSONOutput struct {
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

func (i *Interactor) RenderByJSON(r RenderByJSONInput) RenderByJSONOutput {
	if r.Template == nil || r.Template.Empty() {
		return invalidTemplateBodyOutput
	}

	html := i.Renderer.Render(r.Template)
	return RenderByJSONOutput{Status: StatusSuccess, HTML: html}
}

var invalidTemplateBodyOutput = RenderByJSONOutput{
	Status: StatusInvalid,
	Errors: []templatemanager.ValidationError{
		templatemanager.ValidationError{
			Field:   "template",
			Type:    templatemanager.ErrorInvalid,
			Message: "The given template JSON is invalid",
		},
	},
}
