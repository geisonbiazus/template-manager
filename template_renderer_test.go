package templatemanager

import (
	"testing"

	"github.com/geisonbiazus/addrvrf/assert"
)

func TestTemplateRenderer(t *testing.T) {
	t.Run("Rendering a simple template", func(t *testing.T) {
		renderer := NewTemplateRenderer()
		result := renderer.Render(singleComponent)
		assert.Equal(t, simpleTemplateHTML, result)
	})

	t.Run("Rendering a nested template", func(t *testing.T) {
		renderer := NewTemplateRenderer()
		result := renderer.Render(nestedComponents)
		assert.Equal(t, nestedTemplateHTML, result)
	})
}

var singleComponent = Component{Type: "Page"}
var nestedComponents = Component{
	Type: "Page",
	Children: []Component{
		Component{
			Type: "Section",
			Children: []Component{
				Component{Type: "Text"},
			},
		},
	},
}

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
