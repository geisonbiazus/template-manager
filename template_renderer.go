package templatemanager

import (
	"bytes"
	"strings"
	"text/template"
)

type TemplateRenderer struct {
	tmpl *template.Template
}

func init() {
}

func NewTemplateRenderer() *TemplateRenderer {
	r := &TemplateRenderer{}
	r.loadTemplates()
	return r
}

func (r *TemplateRenderer) Render(c *Component) string {
	buffer := &bytes.Buffer{}
	r.tmpl.ExecuteTemplate(buffer, strings.ToLower(c.Type)+".gohtml", c)
	return buffer.String()
}

func (r *TemplateRenderer) loadTemplates() {
	r.tmpl = template.Must(
		template.New("").Funcs(r.funcMap()).ParseGlob("templates/*"),
	)
}

func (r *TemplateRenderer) funcMap() template.FuncMap {
	return template.FuncMap{
		"Render": r.Render,
	}
}
