package templatemanager

import (
	"testing"

	"github.com/geisonbiazus/templatemanager/assert"
)

func TestRenderTemplateInteractor(t *testing.T) {
	type fixture struct {
		renderer   *RendererSpy
		interactor *RenderTemplateInteractor
		presenter  *TemplatePresenterSpy
	}

	setup := func() *fixture {
		renderer := NewRendererSpy()
		interactor := NewRenderTemplateInteractor(renderer)
		presenter := NewTemplatePresenterSpy()

		return &fixture{
			renderer:   renderer,
			interactor: interactor,
			presenter:  presenter,
		}
	}

	t.Run("RenderByJSON", func(t *testing.T) {
		t.Run("Render and present a valid JSON", func(t *testing.T) {
			f := setup()
			f.renderer.Configure(renderedHTML)

			f.interactor.RenderByJSON(validTemplateJSON, f.presenter)

			assert.DeepEqual(t, validTemplateComponent, f.renderer.Component)
			assert.Equal(t, renderedHTML, f.presenter.HTML)
		})

		t.Run("With invalid json present a validation error", func(t *testing.T) {
			f := setup()
			f.interactor.RenderByJSON(invalidJSON, f.presenter)
			assertInvalidJSONResonse(t, f.presenter)
		})

		t.Run("With empty json present a validation error", func(t *testing.T) {
			f := setup()
			f.interactor.RenderByJSON(emptyJSON, f.presenter)
			assertInvalidJSONResonse(t, f.presenter)
		})
	})
}

func assertInvalidJSONResonse(t *testing.T, p *TemplatePresenterSpy) {
	errors := []ValidationError{
		ValidationError{
			Field:   "template_json",
			Type:    ErrorInvalid,
			Message: "The given template JSON is invalid",
		},
	}

	assert.DeepEqual(t, errors, p.ValidationErrors)
	assert.False(t, p.PresentHTMLCalled)
}

const invalidJSON = `INVALID JSON`
const emptyJSON = `{}`

const validTemplateJSON = `
{
	"type": "Page",
	"children": [
		{
			"type": "Section",
			"children": [
				{
					"type": "Text",
					"properties": {
						"content": "<p>Text</p>"
					}
				}
			]
		}
	]
}
`

var validTemplateComponent = &Component{
	Type: "Page",
	Children: []*Component{
		&Component{
			Type: "Section",
			Children: []*Component{
				&Component{
					Type: "Text",
					Properties: Properties{
						"content": "<p>Text</p>",
					},
				},
			},
		},
	},
}

const renderedHTML = `
<html>
<body>
<section>Text</section>
</body>
</html>
`

type TemplatePresenterSpy struct {
	PresentHTMLCalled bool
	HTML              string
	ValidationErrors  []ValidationError
}

func NewTemplatePresenterSpy() *TemplatePresenterSpy {
	return &TemplatePresenterSpy{}
}

func (p *TemplatePresenterSpy) PresentHTML(html string) {
	p.PresentHTMLCalled = true
	p.HTML = html
}

func (p *TemplatePresenterSpy) PresentValidationErrors(ee []ValidationError) {
	p.ValidationErrors = ee
}

type RendererSpy struct {
	Component *Component
	HTML      string
}

func NewRendererSpy() *RendererSpy {
	return &RendererSpy{}
}

func (r *RendererSpy) Configure(html string) {
	r.HTML = html
}

func (r *RendererSpy) Render(c *Component) string {
	r.Component = c
	return r.HTML
}
