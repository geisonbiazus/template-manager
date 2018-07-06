package rendertemplate_test

import (
	"testing"

	"github.com/geisonbiazus/templatemanager/internal/support/assert"
	"github.com/geisonbiazus/templatemanager/internal/templatemanager"
	. "github.com/geisonbiazus/templatemanager/internal/templatemanager/rendertemplate"
)

func TestInteractor(t *testing.T) {
	type fixture struct {
		renderer   *RendererSpy
		interactor *Interactor
	}

	setup := func() *fixture {
		renderer := NewRendererSpy()
		interactor := NewInteractor(renderer)
		return &fixture{
			renderer:   renderer,
			interactor: interactor,
		}
	}

	t.Run("RenderByJSON", func(t *testing.T) {
		t.Run("Render a valid Component", func(t *testing.T) {
			f := setup()
			f.renderer.Configure(renderedHTML)

			req := RenderByJSONRequest{Template: validComponent}
			resp := f.interactor.RenderByJSON(req)

			assert.Equal(t, req.Template, f.renderer.Component)
			expectedResp := RenderByJSONResponse{Status: StatusSuccess, HTML: renderedHTML}
			assert.DeepEqual(t, expectedResp, resp)
		})

		t.Run("Return a validation error with nil template", func(t *testing.T) {
			f := setup()

			req := RenderByJSONRequest{Template: nil}
			resp := f.interactor.RenderByJSON(req)
			asserInvalidBodyResponse(t, resp)
		})

		t.Run("Return a validation error with empty template", func(t *testing.T) {
			f := setup()
			req := RenderByJSONRequest{Template: &templatemanager.Component{}}
			resp := f.interactor.RenderByJSON(req)
			asserInvalidBodyResponse(t, resp)
		})
	})
}

func asserInvalidBodyResponse(t *testing.T, resp RenderByJSONResponse) {
	t.Helper()
	expected := RenderByJSONResponse{
		Status: StatusInvalid,
		Errors: []templatemanager.ValidationError{
			templatemanager.ValidationError{
				Field:   "body",
				Type:    templatemanager.ErrorInvalid,
				Message: "The given template JSON is invalid",
			},
		},
	}

	assert.DeepEqual(t, expected, resp)
}

var validComponent = &templatemanager.Component{Type: "Page"}

const renderedHTML = `
<html>
<body>
<section>Text</section>
</body>
</html>
`

type RendererSpy struct {
	Component *templatemanager.Component
	HTML      string
}

func NewRendererSpy() *RendererSpy {
	return &RendererSpy{}
}

func (r *RendererSpy) Configure(html string) {
	r.HTML = html
}

func (r *RendererSpy) Render(c *templatemanager.Component) string {
	r.Component = c
	return r.HTML
}
