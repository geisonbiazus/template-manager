package templatemanager

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/geisonbiazus/templatemanager/assert"
)

func TestRenderTemplateByJSONHandler(t *testing.T) {
	type fixture struct {
		handler          *RenderTemplateByJSONHandler
		recorder         *httptest.ResponseRecorder
		interactor       *RenderTemplateInputBoundarySpy
		presenterFactory *RenderTemplateOutputBoundaryFactorySpy
		presenter        *RenderTemplateOutputBoundarySpy
	}

	setup := func() *fixture {
		recorder := httptest.NewRecorder()
		interactor := NewRenderTemplateInputBoundarySpy()
		presenter := NewRenderTemplateOutputBoundarySpy()
		presenterFactory := NewRenderTemplateOutputBoundaryFactorySpy()
		presenterFactory.Configure(presenter)
		handler := NewRenderTemplateByJSONHandler(interactor, presenterFactory)

		return &fixture{
			handler:          handler,
			recorder:         recorder,
			interactor:       interactor,
			presenterFactory: presenterFactory,
			presenter:        presenter,
		}
	}

	const validRequestBody = `{"template": {"type":"Page"}}`

	t.Run("template JSON goes through interactor", func(t *testing.T) {
		f := setup()
		r := httptest.NewRequest(http.MethodPost, "http://example.org", bytes.NewBufferString(validRequestBody))

		f.handler.ServeHTTP(f.recorder, r)

		assert.DeepEqual(t, &Component{Type: "Page"}, f.interactor.Template)
		assert.Equal(t, f.presenter, f.interactor.Presenter)
		assert.Equal(t, f.presenterFactory.ResponseWriter, f.recorder)
	})
}

type RenderTemplateInputBoundarySpy struct {
	Template  *Component
	Presenter RenderTemplateOutputBoundary
}

func NewRenderTemplateInputBoundarySpy() *RenderTemplateInputBoundarySpy {
	return &RenderTemplateInputBoundarySpy{}
}

func (r *RenderTemplateInputBoundarySpy) RenderByJSON(
	template *Component, presenter RenderTemplateOutputBoundary,
) {
	r.Template = template
	r.Presenter = presenter
}

type RenderTemplateOutputBoundaryFactorySpy struct {
	ResponseWriter http.ResponseWriter
	Presenter      *RenderTemplateOutputBoundarySpy
}

func NewRenderTemplateOutputBoundaryFactorySpy() *RenderTemplateOutputBoundaryFactorySpy {
	return &RenderTemplateOutputBoundaryFactorySpy{}
}

func (f *RenderTemplateOutputBoundaryFactorySpy) Create(w http.ResponseWriter) RenderTemplateOutputBoundary {
	f.ResponseWriter = w
	return f.Presenter
}

func (f *RenderTemplateOutputBoundaryFactorySpy) Configure(p *RenderTemplateOutputBoundarySpy) {
	f.Presenter = p
}
