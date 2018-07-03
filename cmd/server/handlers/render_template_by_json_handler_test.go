package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/geisonbiazus/templatemanager/internal/support/assert"
	"github.com/geisonbiazus/templatemanager/internal/templatemanager"
)

func TestRenderTemplateByJSONHandler(t *testing.T) {
	type fixture struct {
		handler       *RenderTemplateByJSONHandler
		recorder      *httptest.ResponseRecorder
		input         *RenderTemplateInputBoundarySpy
		outputFactory *RenderTemplateOutputBoundaryFactorySpy
		output        *RenderTemplateOutputBoundarySpy
	}

	setup := func() *fixture {
		recorder := httptest.NewRecorder()
		input := NewRenderTemplateInputBoundarySpy()
		output := NewRenderTemplateOutputBoundarySpy()
		outputFactory := NewRenderTemplateOutputBoundaryFactorySpy()
		outputFactory.Configure(output)
		handler := NewRenderTemplateByJSONHandler(input, outputFactory)

		return &fixture{
			handler:       handler,
			recorder:      recorder,
			input:         input,
			outputFactory: outputFactory,
			output:        output,
		}
	}

	const validRequestBody = `{"template": {"body": {"type":"Page"}}}`

	t.Run("template JSON goes through input", func(t *testing.T) {
		f := setup()
		r := httptest.NewRequest(http.MethodPost, "http://example.org", bytes.NewBufferString(validRequestBody))

		f.handler.ServeHTTP(f.recorder, r)

		assert.DeepEqual(t, &templatemanager.Component{Type: "Page"}, f.input.Template)
		assert.Equal(t, f.output, f.input.Output)
		assert.Equal(t, f.outputFactory.ResponseWriter, f.recorder)
	})
}

type RenderTemplateInputBoundarySpy struct {
	Template *templatemanager.Component
	Output   templatemanager.RenderTemplateOutputBoundary
}

func NewRenderTemplateInputBoundarySpy() *RenderTemplateInputBoundarySpy {
	return &RenderTemplateInputBoundarySpy{}
}

func (r *RenderTemplateInputBoundarySpy) RenderByJSON(
	template *templatemanager.Component, output templatemanager.RenderTemplateOutputBoundary,
) {
	r.Template = template
	r.Output = output
}

type RenderTemplateOutputBoundaryFactorySpy struct {
	ResponseWriter http.ResponseWriter
	Output         *RenderTemplateOutputBoundarySpy
}

func NewRenderTemplateOutputBoundaryFactorySpy() *RenderTemplateOutputBoundaryFactorySpy {
	return &RenderTemplateOutputBoundaryFactorySpy{}
}

func (f *RenderTemplateOutputBoundaryFactorySpy) Create(w http.ResponseWriter) templatemanager.RenderTemplateOutputBoundary {
	f.ResponseWriter = w
	return f.Output
}

func (f *RenderTemplateOutputBoundaryFactorySpy) Configure(output *RenderTemplateOutputBoundarySpy) {
	f.Output = output
}

type RenderTemplateOutputBoundarySpy struct {
	PresentHTMLCalled bool
	HTML              string
	ValidationErrors  []templatemanager.ValidationError
}

func NewRenderTemplateOutputBoundarySpy() *RenderTemplateOutputBoundarySpy {
	return &RenderTemplateOutputBoundarySpy{}
}

func (p *RenderTemplateOutputBoundarySpy) PresentHTML(html string) {}
func (p *RenderTemplateOutputBoundarySpy) PresentValidationErrors(ee []templatemanager.ValidationError) {
}
