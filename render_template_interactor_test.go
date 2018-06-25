package templatemanager

import (
	"testing"

	"github.com/geisonbiazus/templatemanager/assert"
)

func TestRenderTemplateInteractor(t *testing.T) {
	t.Run("RenderByJSON", func(t *testing.T) {
		t.Run("Renders simple template", func(t *testing.T) {
			interactor := NewRenderTemplateInteractor()
			result := interactor.RenderByJSON(simplePageJSON)
			assert.Equal(t, simpleTemplateHTML, result)
		})

		t.Run("Renders a nested template", func(t *testing.T) {
			interactor := NewRenderTemplateInteractor()
			result := interactor.RenderByJSON(nestedTemplateJSON)
			assert.Equal(t, nestedTemplateHTML, result)
		})
	})
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
