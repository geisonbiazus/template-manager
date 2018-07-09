package rendertemplate_test

import (
	"testing"

	"github.com/geisonbiazus/templatemanager/internal/support/assert"
	"github.com/geisonbiazus/templatemanager/internal/templatemanager"
	. "github.com/geisonbiazus/templatemanager/internal/templatemanager/rendertemplate"
)

func TestInteractor(t *testing.T) {
	type fixture struct {
		renderer   *RendererSpy
		interactor *Interactor
	}

	setup := func() *fixture {
		renderer := NewRendererSpy()
		interactor := NewInteractor(renderer)
		return &fixture{
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

			assert.Equal(t, input.Template, f.renderer.Component)
			expectedResp := RenderByJSONOutput{Status: StatusSuccess, HTML: renderedHTML}
			assert.DeepEqual(t, expectedResp, output)
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
}

func asserInvalidBodyResponse(t *testing.T, output RenderByJSONOutput) {
	t.Helper()
	expected := RenderByJSONOutput{
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
