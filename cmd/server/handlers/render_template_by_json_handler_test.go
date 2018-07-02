package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/geisonbiazus/templatemanager/cmd/server/presenters"
	"github.com/geisonbiazus/templatemanager/pkg/support/assert"
	"github.com/geisonbiazus/templatemanager/pkg/templatemanager"
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

	const validRequestBody = `{"template": {"type":"Page"}}`

	t.Run("template JSON goes through input", func(t *testing.T) {
		f := setup()
		r := httptest.NewRequest(http.MethodPost, "http://example.org", bytes.NewBufferString(validRequestBody))

		f.handler.ServeHTTP(f.recorder, r)

		assert.DeepEqual(t, &templatemanager.Component{Type: "Page"}, f.input.Template)
		assert.Equal(t, f.output, f.input.Output)
		assert.Equal(t, f.outputFactory.ResponseWriter, f.recorder)
	})
}

func TestRenderTemplateByJSONIntegration(t *testing.T) {
	renderer := templatemanager.NewTemplateRenderer("../../../" + templatemanager.DefaultTemplatePath)
	input := templatemanager.NewRenderTemplateInteractor(renderer)
	outputFactory := presenters.NewRenderTemplateJSONPresenterFactory()
	handler := NewRenderTemplateByJSONHandler(input, outputFactory)

	body := bytes.NewBufferString(`{"template": {"type":"Page"}}`)
	r := httptest.NewRequest(http.MethodPost, "http://example.org", body)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, r)

	expected := `{"html":"\u003c!DOCTYPE html\u003e\n\u003chtml\u003e\n\u003chead\u003e\n\u003cmeta charset=\"UTF-8\"\u003e\n\u003ctitle\u003e\u003c/title\u003e\n\u003c/head\u003e\n\u003cbody\u003e\n  \n\u003c/body\u003e\n\u003c/html\u003e\n"}` + "\n"
	response := w.Body.String()
	assert.Equal(t, expected, response)
	assert.Equal(t, http.StatusOK, w.Code)
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
