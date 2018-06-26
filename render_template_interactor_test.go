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
	})
}

type TemplatePresenterSpy struct {
	HTML string
}

func NewTemplatePresenterSpy() *TemplatePresenterSpy {
	return &TemplatePresenterSpy{}
}

func (p *TemplatePresenterSpy) PresentHTML(html string) {
	p.HTML = html
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

const validTemplateJSON = `
{
	"type": "Page",
	"children": [
		{
			"type": "Section",
			"children": [
				{
					"type": "Text"
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
				&Component{Type: "Text"},
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
