package handlers_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/geisonbiazus/templatemanager/cmd/server/handlers"
	"github.com/geisonbiazus/templatemanager/internal/support/assert"
	"github.com/geisonbiazus/templatemanager/internal/templatemanager"
)

func TestCreateTemplateHandler(t *testing.T) {
	input := NewManageTemplateInputBoundarySpy()
	output := NewManageTemplateOutputBoundarySpy()
	outputFactory := NewManageTemplateOutputBoundaryFactorySpy()
	outputFactory.Output = output
	handler := NewCreateTemplateHandler(input, outputFactory)
	w := httptest.NewRecorder()

	body := bytes.NewBufferString(`{"template": {"body": {"type": "Page"}}}`)

	r := httptest.NewRequest(http.MethodPost, "http://example.org", body)
	handler.ServeHTTP(w, r)

	expectedTemplate := templatemanager.Template{
		Body: &templatemanager.Component{
			Type: "Page",
		},
	}
	assert.DeepEqual(t, expectedTemplate, input.Template)
	assert.Equal(t, output, input.Output)
	assert.Equal(t, w, outputFactory.Writer)
}

type ManageTemplateInputBoundarySpy struct {
	Template templatemanager.Template
	Output   templatemanager.ManageTemplateOutputBoundary
}

func NewManageTemplateInputBoundarySpy() *ManageTemplateInputBoundarySpy {
	return &ManageTemplateInputBoundarySpy{}
}

func (i *ManageTemplateInputBoundarySpy) Create(
	t templatemanager.Template,
	output templatemanager.ManageTemplateOutputBoundary,
) {
	i.Template = t
	i.Output = output
}

type ManageTemplateOutputBoundaryFactorySpy struct {
	Writer http.ResponseWriter
	Output *ManageTemplateOutputBoundarySpy
}

func NewManageTemplateOutputBoundaryFactorySpy() *ManageTemplateOutputBoundaryFactorySpy {
	return &ManageTemplateOutputBoundaryFactorySpy{}
}

func (f *ManageTemplateOutputBoundaryFactorySpy) Create(
	w http.ResponseWriter,
) templatemanager.ManageTemplateOutputBoundary {
	f.Writer = w
	return f.Output
}

type ManageTemplateOutputBoundarySpy struct{}

func NewManageTemplateOutputBoundarySpy() *ManageTemplateOutputBoundarySpy {
	return &ManageTemplateOutputBoundarySpy{}
}

func (o *ManageTemplateOutputBoundarySpy) PresentCreated(template templatemanager.Template)          {}
func (o *ManageTemplateOutputBoundarySpy) PresentValidationErrors([]templatemanager.ValidationError) {}
