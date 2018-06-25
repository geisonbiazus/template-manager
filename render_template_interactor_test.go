package templatemanager

import (
	"testing"

	"github.com/geisonbiazus/templatemanager/assert"
)

func TestRenderTemplateInteractor(t *testing.T) {
	t.Run("RenderByJSON", func(t *testing.T) {
		t.Run("Renders simple tempalte JSON", func(t *testing.T) {
			interactor := NewRenderTemplateInteractor()
			result := interactor.RenderByJSON(simplePageJSON)
			assert.Equal(t, simplePageHTML, result)
		})
	})
}

const simplePageJSON = `
{
	"type": "Page",
}
`

const simplePageHTML = `<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<title></title>
</head>
<body>
</body>
</html>
`
