package template_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/geisonbiazus/templatemanager/cmd/server/handlers/template"
	"github.com/geisonbiazus/templatemanager/internal/support/assert"
	"github.com/geisonbiazus/templatemanager/internal/templatemanager"
	"github.com/geisonbiazus/templatemanager/internal/templatemanager/rendertemplate"
)

func TestRenderByJSONHandler(t *testing.T) {
	type fixture struct {
		interactor *RenderTemplateInteractorSpy
		handler    *RenderByJSONHandler
		writer     *httptest.ResponseRecorder
	}

	setup := func() *fixture {
		interactor := NewRenderTemplateInteractorSpy()
		handler := NewRenderByJSONHandler(interactor)
		w := httptest.NewRecorder()
		return &fixture{
			interactor: interactor,
			handler:    handler,
			writer:     w,
		}
	}

	t.Run("Successful rendering", func(t *testing.T) {
		f := setup()

		r := newRequest(`{"template": {"body": {"type":"Page"}}}`)
		f.interactor.ConfigureRenderByJSONSuccessResponse("rendered html")

		f.handler.ServeHTTP(f.writer, r)

		assertRenderByJSONSuccessResponse(t, f.interactor, f.writer)
	})
}

func newRequest(body string) *http.Request {
	buffer := bytes.NewBufferString(body)

	return httptest.NewRequest(http.MethodPost, "http://example.org/", buffer)
}

func assertRenderByJSONSuccessResponse(
	t *testing.T, i *RenderTemplateInteractorSpy, w *httptest.ResponseRecorder,
) {
	expectedReq := rendertemplate.RenderByJSONRequest{
		Template: rendertemplate.Template{
			Body: &templatemanager.Component{Type: "Page"},
		},
	}

	assert.DeepEqual(t, expectedReq, i.Request)

	expected := `{"html":"rendered html"}` + "\n"
	assert.Equal(t, expected, w.Body.String())
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
}

type RenderTemplateInteractorSpy struct {
	Request  rendertemplate.RenderByJSONRequest
	Response rendertemplate.RenderByJSONResponse
}

func NewRenderTemplateInteractorSpy() *RenderTemplateInteractorSpy {
	return &RenderTemplateInteractorSpy{}
}

func (i *RenderTemplateInteractorSpy) RenderByJSON(
	r rendertemplate.RenderByJSONRequest,
) rendertemplate.RenderByJSONResponse {
	i.Request = r
	return i.Response
}

func (i *RenderTemplateInteractorSpy) ConfigureRenderByJSONSuccessResponse(html string) {
	i.Response = rendertemplate.RenderByJSONResponse{
		Status: rendertemplate.StatusSuccess,
		HTML:   html,
	}
}
