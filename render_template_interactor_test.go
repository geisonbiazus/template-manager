package templatemanager

import (
	"testing"

	"github.com/geisonbiazus/templatemanager/assert"
)

func TestRenderTemplateInteractor(t *testing.T) {
	type fixture struct {
		renderer   *RendererSpy
		interactor *RenderTemplateInteractor
		output     *RenderTemplateOutputBoundarySpy
	}

	setup := func() *fixture {
		renderer := NewRendererSpy()
		interactor := NewRenderTemplateInteractor(renderer)
		output := NewRenderTemplateOutputBoundarySpy()

		return &fixture{
			renderer:   renderer,
			interactor: interactor,
			output:     output,
		}
	}

	t.Run("RenderByJSON", func(t *testing.T) {
		t.Run("Render and present a valid Component", func(t *testing.T) {
			f := setup()
			f.renderer.Configure(renderedHTML)

			template := &Component{Type: "Page"}

			f.interactor.RenderByJSON(template, f.output)

			assert.DeepEqual(t, template, f.renderer.Component)
			assert.Equal(t, renderedHTML, f.output.HTML)
		})

		t.Run("Present a validation error with nil template", func(t *testing.T) {
			f := setup()
			f.interactor.RenderByJSON(nil, f.output)
			assertInvalidJSONResonse(t, f.output)
		})

		t.Run("Present a validation error with empty template", func(t *testing.T) {
			f := setup()
			f.interactor.RenderByJSON(&Component{}, f.output)
			assertInvalidJSONResonse(t, f.output)
		})
	})
}

func assertInvalidJSONResonse(t *testing.T, p *RenderTemplateOutputBoundarySpy) {
	t.Helper()
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

const renderedHTML = `
<html>
<body>
<section>Text</section>
</body>
</html>
`

type RenderTemplateOutputBoundarySpy struct {
	PresentHTMLCalled bool
	HTML              string
	ValidationErrors  []ValidationError
}

func NewRenderTemplateOutputBoundarySpy() *RenderTemplateOutputBoundarySpy {
	return &RenderTemplateOutputBoundarySpy{}
}

func (p *RenderTemplateOutputBoundarySpy) PresentHTML(html string) {
	p.PresentHTMLCalled = true
	p.HTML = html
}

func (p *RenderTemplateOutputBoundarySpy) PresentValidationErrors(ee []ValidationError) {
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
