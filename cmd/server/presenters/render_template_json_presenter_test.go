package presenters

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/geisonbiazus/templatemanager/assert"
	"github.com/geisonbiazus/templatemanager/pkg/templatemanager"
)

func TestRenderTemplateJSONPresenter(t *testing.T) {
	type fixture struct {
		presenter *RenderTemplateJSONPresenter
		writer    *httptest.ResponseRecorder
	}

	setup := func() *fixture {
		factory := NewRenderTemplateJSONPresenterFactory()
		w := httptest.NewRecorder()
		presenter := factory.Create(w)

		return &fixture{
			presenter: presenter.(*RenderTemplateJSONPresenter),
			writer:    w,
		}
	}

	t.Run("PresentHTML", func(t *testing.T) {
		t.Run("Presents the success result", func(t *testing.T) {
			f := setup()
			html := "<html></html>"
			f.presenter.PresentHTML(html)

			expected := `{"html":"\u003chtml\u003e\u003c/html\u003e"}` + "\n"
			assert.Equal(t, expected, f.writer.Body.String())
			assert.Equal(t, http.StatusOK, f.writer.Code)
			assert.Equal(t, "application/json", f.writer.Header().Get("Content-Type"))
		})
	})

	t.Run("PresentValidationErrors", func(t *testing.T) {
		t.Run("Presents the errors as json", func(t *testing.T) {
			f := setup()
			errors := []templatemanager.ValidationError{
				templatemanager.ValidationError{
					Field:   "field",
					Message: "message",
					Type:    "type",
				},
			}
			f.presenter.PresentValidationErrors(errors)

			expected := `{"errors":[{"field":"field","type":"type","message":"message"}]}` + "\n"

			assert.Equal(t, expected, f.writer.Body.String())
			assert.Equal(t, http.StatusUnprocessableEntity, f.writer.Code)
			assert.Equal(t, "application/json", f.writer.Header().Get("Content-Type"))
		})
	})
}
