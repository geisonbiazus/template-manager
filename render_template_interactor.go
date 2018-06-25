package templatemanager

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"
)

type RenderTemplateInteractor struct {
}

func NewRenderTemplateInteractor() *RenderTemplateInteractor {
	return &RenderTemplateInteractor{}
}

var fm = template.FuncMap{
	"renderComponent": renderComponent,
}

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.New("").Funcs(fm).ParseGlob("templates/*"))
}

func (r *RenderTemplateInteractor) RenderByJSON(templateJSON string) string {
	component := Component{}
	json.Unmarshal([]byte(templateJSON), &component)

	return renderComponent(component)
}

func renderComponent(c Component) string {
	buffer := &bytes.Buffer{}

	tmpl.ExecuteTemplate(buffer, strings.ToLower(c.Type)+".gohtml", c)

	return buffer.String()
}

type Component struct {
	Type     string      `json:"type"`
	Children []Component `json:"children"`
}
