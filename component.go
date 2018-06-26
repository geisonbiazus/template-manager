package templatemanager

import (
	"bytes"
	"strings"
	"text/template"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*"))
}

type Component struct {
	Type     string       `json:"type"`
	Children []*Component `json:"children"`
}

func (c *Component) Empty() bool {
	return c.Type == "" && len(c.Children) == 0
}

func (c *Component) Render() string {
	buffer := &bytes.Buffer{}

	tmpl.ExecuteTemplate(buffer, strings.ToLower(c.Type)+".gohtml", c)

	return buffer.String()
}

const (
	ErrorInvalid = "INVALID"
)

type ValidationError struct {
	Field   string
	Type    string
	Message string
}
