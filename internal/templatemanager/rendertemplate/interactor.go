package rendertemplate

import (
	"github.com/geisonbiazus/templatemanager/internal/templatemanager"
)

const (
	StatusSuccess = "Success"
	StatusInvalid = "Invalid"
)

type TemplateRespository interface {
	FindByID(id string) (templatemanager.Template, error)
}

type Renderer interface {
	Render(*templatemanager.Component) string
}

type RenderByJSONInput struct {
	Template *templatemanager.Component
}

type RenderByIDInput struct {
	ID string
}

type Output struct {
	Status string
	HTML   string
	Errors []templatemanager.ValidationError
}

type Interactor struct {
	Renderer   Renderer
	Repository TemplateRespository
}

func NewInteractor(renderer Renderer, repository TemplateRespository) *Interactor {
	return &Interactor{
		Renderer:   renderer,
		Repository: repository,
	}
}

func (i *Interactor) RenderByJSON(input RenderByJSONInput) Output {
	if input.Template == nil || input.Template.Empty() {
		return invalidTemplateBodyOutput
	}
	return i.renderAndCreateOutput(input.Template)
}

func (i *Interactor) RenderByID(input RenderByIDInput) Output {
	template, _ := i.Repository.FindByID(input.ID)
	return i.renderAndCreateOutput(template.Component)
}

func (i *Interactor) renderAndCreateOutput(c *templatemanager.Component) Output {
	html := i.Renderer.Render(c)
	return Output{Status: StatusSuccess, HTML: html}
}

var invalidTemplateBodyOutput = Output{
	Status: StatusInvalid,
	Errors: []templatemanager.ValidationError{
		templatemanager.ValidationError{
			Field:   "template",
			Type:    templatemanager.ErrorInvalid,
			Message: "The given template JSON is invalid",
		},
	},
}
