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
		t.Run("Render and present a valid Component", func(t *testing.T) {
			f := setup()
			f.renderer.Configure(renderedHTML)

			template := &Component{Type: "Page"}

			f.interactor.RenderByJSON(template, f.presenter)

			assert.DeepEqual(t, template, f.renderer.Component)
			assert.Equal(t, renderedHTML, f.presenter.HTML)
		})

		t.Run("Present a validation error with nil template", func(t *testing.T) {
			f := setup()
			f.interactor.RenderByJSON(nil, f.presenter)
			assertInvalidJSONResonse(t, f.presenter)
		})

		t.Run("Present a validation error with empty template", func(t *testing.T) {
			f := setup()
			f.interactor.RenderByJSON(&Component{}, f.presenter)
			assertInvalidJSONResonse(t, f.presenter)
		})
	})
}

func assertInvalidJSONResonse(t *testing.T, p *TemplatePresenterSpy) {
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
