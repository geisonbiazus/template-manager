package templatemanager

import (
	"bytes"
	"html/template"
)

type RenderTemplateInteractor struct {
}

func NewRenderTemplateInteractor() *RenderTemplateInteractor {
	return &RenderTemplateInteractor{}
}

func (i *RenderTemplateInteractor) RenderByJSON(templateJSON string) string {
	tmpl := template.Must(template.ParseGlob("templates/page.gohtml"))

	buffer := &bytes.Buffer{}
	tmpl.Execute(buffer, nil)

	return buffer.String()
}
