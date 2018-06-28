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
		handler    *RenderTemplateByJSONHandler
		recorder   *httptest.ResponseRecorder
		interactor *RendererInteractorSpy
		presenter  *HTTPRenderTemplatePresenterSpy
	}

	setup := func() *fixture {
		recorder := httptest.NewRecorder()
		interactor := NewRendererInteractorSpy()
		presenter := NewHTTPRenderTemplatePresenterSpy()
		handler := NewRenderTemplateByJSONHandler(interactor, presenter)

		return &fixture{
			handler:    handler,
			recorder:   recorder,
			interactor: interactor,
			presenter:  presenter,
		}
	}

	const validRequestBody = `{"template": {"type":"Page"}}`

	t.Run("template JSON goes through interactor", func(t *testing.T) {
		f := setup()
		r := httptest.NewRequest(http.MethodPost, "http://example.org", bytes.NewBufferString(validRequestBody))
		f.handler.ServeHTTP(f.recorder, r)
		assert.DeepEqual(t, &Component{Type: "Page"}, f.interactor.Template)
		assert.Equal(t, f.presenter, f.interactor.Presenter)
		assert.Equal(t, f.presenter.ResponseWriter, f.recorder)
	})
}

type RendererInteractorSpy struct {
	Template  *Component
	Presenter RenderTemplatePresenter
}

func NewRendererInteractorSpy() *RendererInteractorSpy {
	return &RendererInteractorSpy{}
}

func (r *RendererInteractorSpy) RenderByJSON(
	template *Component, presenter RenderTemplatePresenter,
) {
	r.Template = template
	r.Presenter = presenter
}

type HTTPRenderTemplatePresenterSpy struct {
	ResponseWriter http.ResponseWriter
}

func NewHTTPRenderTemplatePresenterSpy() *HTTPRenderTemplatePresenterSpy {
	return &HTTPRenderTemplatePresenterSpy{}
}

func (p *HTTPRenderTemplatePresenterSpy) With(w http.ResponseWriter) HTTPRenderTemplatePresenter {
	p.ResponseWriter = w
	return p
}

func (p *HTTPRenderTemplatePresenterSpy) PresentHTML(html string)                      {}
func (p *HTTPRenderTemplatePresenterSpy) PresentValidationErrors(ee []ValidationError) {}
