package templatemanager

import (
	"bytes"
	"strings"
	"text/template"
)

const DefaultTemplatePath = "internal/templatemanager/test/templates/*"

type TemplateRenderer struct {
	tmpl *template.Template
}

func NewTemplateRenderer(templatesPath string) *TemplateRenderer {
	r := &TemplateRenderer{}
	r.loadTemplates(templatesPath)
	return r
}

func (r *TemplateRenderer) Render(c *Component) string {
	buffer := &bytes.Buffer{}
	r.tmpl.ExecuteTemplate(buffer, strings.ToLower(c.Type)+".gohtml", c)
	return buffer.String()
}

func (r *TemplateRenderer) loadTemplates(templatesPath string) {
	r.tmpl = template.Must(
		template.New("").Funcs(r.funcMap()).ParseGlob(templatesPath),
	)
}

func (r *TemplateRenderer) funcMap() template.FuncMap {
	return template.FuncMap{
		"Render": r.Render,
	}
}
