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
		interactor       *RendererInteractorSpy
		presenterFactory *HTTPRenderTemplatePresenterFactorySpy
		presenter        *RenderTemplatePresenterSpy
	}

	setup := func() *fixture {
		recorder := httptest.NewRecorder()
		interactor := NewRendererInteractorSpy()
		presenter := NewRenderTemplatePresenterSpy()
		presenterFactory := NewHTTPRenderTemplatePresenterFactorySpy()
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

type HTTPRenderTemplatePresenterFactorySpy struct {
	ResponseWriter http.ResponseWriter
	Presenter      *RenderTemplatePresenterSpy
}

func NewHTTPRenderTemplatePresenterFactorySpy() *HTTPRenderTemplatePresenterFactorySpy {
	return &HTTPRenderTemplatePresenterFactorySpy{}
}

func (f *HTTPRenderTemplatePresenterFactorySpy) Create(w http.ResponseWriter) RenderTemplatePresenter {
	f.ResponseWriter = w
	return f.Presenter
}

func (f *HTTPRenderTemplatePresenterFactorySpy) Configure(p *RenderTemplatePresenterSpy) {
	f.Presenter = p
}
