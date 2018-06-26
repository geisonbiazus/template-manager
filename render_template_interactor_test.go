package templatemanager

import (
	"testing"

	"github.com/geisonbiazus/templatemanager/assert"
)

func TestRenderTemplateInteractor(t *testing.T) {
	type fixture struct {
		interactor *RenderTemplateInteractor
		presenter  *TemplatePresenterSpy
	}

	setup := func() *fixture {
		presenter := NewTemplatePresenterSpy()
		interactor := NewRenderTemplateInteractor()

		return &fixture{
			interactor: interactor,
			presenter:  presenter,
		}
	}

	t.Run("RenderByJSON", func(t *testing.T) {
		t.Run("Rendering a simple template", func(t *testing.T) {
			f := setup()
			f.interactor.RenderByJSON(simplePageJSON, f.presenter)
			assert.Equal(t, simpleTemplateHTML, f.presenter.HTML)
		})

		t.Run("Rendering a nested template", func(t *testing.T) {
			f := setup()
			f.interactor.RenderByJSON(nestedTemplateJSON, f.presenter)
			assert.Equal(t, nestedTemplateHTML, f.presenter.HTML)
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

const simplePageJSON = `
{
	"type": "Page"
}
`

const nestedTemplateJSON = `
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

const simpleTemplateHTML = "<!DOCTYPE html>\n" +
	"<html>\n" +
	"<head>\n" +
	"<meta charset=\"UTF-8\">\n" +
	"<title></title>\n" +
	"</head>\n" +
	"<body>\n" +
	"  \n" +
	"</body>\n" +
	"</html>\n"

const nestedTemplateHTML = "<!DOCTYPE html>\n" +
	"<html>\n" +
	"<head>\n" +
	"<meta charset=\"UTF-8\">\n" +
	"<title></title>\n" +
	"</head>\n" +
	"<body>\n" +
	"  \n" +
	"    <section>\n" +
	"  \n" +
	"    <p>Text</p>\n" +
	"\n" +
	"  \n" +
	"</section>\n" +
	"\n" +
	"  \n" +
	"</body>\n" +
	"</html>\n"
