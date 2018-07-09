package rendertemplate_test

import (
	"testing"

	"github.com/geisonbiazus/templatemanager/internal/support/assert"
	"github.com/geisonbiazus/templatemanager/internal/templatemanager"
	. "github.com/geisonbiazus/templatemanager/internal/templatemanager/rendertemplate"
)

func TestInteractor(t *testing.T) {
	type fixture struct {
		repository *templatemanager.InMemoryTemplateRepository
		renderer   *RendererSpy
		interactor *Interactor
	}

	setup := func() *fixture {
		repository := templatemanager.NewInMemoryTemplateRepository()
		renderer := NewRendererSpy()
		interactor := NewInteractor(renderer, repository)
		return &fixture{
			repository: repository,
			renderer:   renderer,
			interactor: interactor,
		}
	}

	t.Run("RenderByJSON", func(t *testing.T) {
		t.Run("Render a valid Component", func(t *testing.T) {
			f := setup()
			f.renderer.Configure(renderedHTML)

			input := RenderByJSONInput{Template: validComponent}
			output := f.interactor.RenderByJSON(input)

			assertSuccessResponse(t, f.renderer, output, renderedHTML, input.Template)
		})

		t.Run("Return a validation error with nil template", func(t *testing.T) {
			f := setup()

			input := RenderByJSONInput{Template: nil}
			output := f.interactor.RenderByJSON(input)
			asserInvalidBodyResponse(t, output)
		})

		t.Run("Return a validation error with empty template", func(t *testing.T) {
			f := setup()
			input := RenderByJSONInput{Template: &templatemanager.Component{}}
			output := f.interactor.RenderByJSON(input)
			asserInvalidBodyResponse(t, output)
		})
	})

	t.Run("RenderByID", func(t *testing.T) {
		f := setup()
		template := addTemplateToRepository(f.repository)
		f.renderer.Configure(renderedHTML)

		input := RenderByIDInput{ID: "1"}
		output := f.interactor.RenderByID(input)

		assertSuccessResponse(t, f.renderer, output, renderedHTML, template.Component)
	})
}

func addTemplateToRepository(r *templatemanager.InMemoryTemplateRepository) templatemanager.Template {
	component := &templatemanager.Component{Type: "Page"}
	template := templatemanager.Template{Component: component}
	r.Create(template)
	return template
}

func assertSuccessResponse(
	t *testing.T,
	renderer *RendererSpy,
	output Output,
	expectedHTML string,
	expectedComponent *templatemanager.Component,
) {
	t.Helper()

	expectedOutput := Output{Status: StatusSuccess, HTML: expectedHTML}
	assert.DeepEqual(t, expectedOutput, output)
	assert.Equal(t, expectedComponent, renderer.Component)
}

func asserInvalidBodyResponse(t *testing.T, output Output) {
	t.Helper()
	expected := Output{
		Status: StatusInvalid,
		Errors: []templatemanager.ValidationError{
			templatemanager.ValidationError{
				Field:   "template",
				Type:    templatemanager.ErrorInvalid,
				Message: "The given template JSON is invalid",
			},
		},
	}

	assert.DeepEqual(t, expected, output)
}

var validComponent = &templatemanager.Component{Type: "Page"}

const renderedHTML = `
<html>
<body>
<section>Text</section>
</body>
</html>
`

type RendererSpy struct {
	Component *templatemanager.Component
	HTML      string
}

func NewRendererSpy() *RendererSpy {
	return &RendererSpy{}
}

func (r *RendererSpy) Configure(html string) {
	r.HTML = html
}

func (r *RendererSpy) Render(c *templatemanager.Component) string {
	r.Component = c
	return r.HTML
}
