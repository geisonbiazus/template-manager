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
		f.interactor.ConfigureRenderByJSONSuccessOutput("rendered html")

		f.handler.ServeHTTP(f.writer, r)

		assertRenderByJSONSuccessResponse(t, f.interactor, f.writer)
	})

	t.Run("Validation error on rendering", func(t *testing.T) {
		f := setup()

		r := newRequest(`{"template": {"body": null}}`)
		f.interactor.ConfigureRenderByJSONInvalidOutput()

		f.handler.ServeHTTP(f.writer, r)

		assertRenderByJSONInvalidResponse(t, f.interactor, f.writer)
	})
}

func newRequest(body string) *http.Request {
	buffer := bytes.NewBufferString(body)

	return httptest.NewRequest(http.MethodPost, "http://example.org/", buffer)
}

func assertRenderByJSONSuccessResponse(
	t *testing.T, i *RenderTemplateInteractorSpy, w *httptest.ResponseRecorder,
) {
	expectedInput := rendertemplate.RenderByJSONInput{
		Template: &templatemanager.Component{Type: "Page"},
	}

	assert.DeepEqual(t, expectedInput, i.RenderByJSONInput)

	expected := `{"html":"rendered html"}` + "\n"
	assert.Equal(t, expected, w.Body.String())
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
}

func assertRenderByJSONInvalidResponse(
	t *testing.T, i *RenderTemplateInteractorSpy, w *httptest.ResponseRecorder,
) {
	expected := `{"errors":[{"field":"FIELD","type":"TYPE","message":"MESSAGE"}]}` + "\n"
	assert.Equal(t, expected, w.Body.String())
	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
}

type RenderTemplateInteractorSpy struct {
	RenderByJSONInput  rendertemplate.RenderByJSONInput
	RenderByJSONOutput rendertemplate.RenderByJSONOutput
}

func NewRenderTemplateInteractorSpy() *RenderTemplateInteractorSpy {
	return &RenderTemplateInteractorSpy{}
}

func (i *RenderTemplateInteractorSpy) RenderByJSON(
	input rendertemplate.RenderByJSONInput,
) rendertemplate.RenderByJSONOutput {
	i.RenderByJSONInput = input
	return i.RenderByJSONOutput
}

func (i *RenderTemplateInteractorSpy) ConfigureRenderByJSONSuccessOutput(html string) {
	i.RenderByJSONOutput = rendertemplate.RenderByJSONOutput{
		Status: rendertemplate.StatusSuccess,
		HTML:   html,
	}
}

func (i *RenderTemplateInteractorSpy) ConfigureRenderByJSONInvalidOutput() {
	i.RenderByJSONOutput = rendertemplate.RenderByJSONOutput{
		Status: rendertemplate.StatusInvalid,
		Errors: []templatemanager.ValidationError{
			templatemanager.ValidationError{
				Type:    "TYPE",
				Field:   "FIELD",
				Message: "MESSAGE",
			},
		},
	}
}
