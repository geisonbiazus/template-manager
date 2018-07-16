package templatemanager

import (
	"testing"

	"github.com/geisonbiazus/templatemanager/internal/support/assert"
)

type RenderPayloadFixture struct {
	renderer  *RendererSpy
	presenter *ContentPresenterSpy
}

func TestRenderPayloadInteractor(t *testing.T) {
	setup := func() *RenderPayloadFixture {
		return &RenderPayloadFixture{
			renderer:  NewRendererSpy(),
			presenter: NewContentPresenterSpy(),
		}
	}

	t.Run("It renders a payload and presents it", func(t *testing.T) {
		f := setup()

		payload := &Component{Type: "Page"}
		expectedContent := "RenderedContent"
		f.renderer.Content = expectedContent

		service := NewRenderPayloadService(payload, f.renderer, f.presenter)

		service.Execute()

		assert.DeepEqual(t, payload, f.renderer.Component)
		assert.Equal(t, expectedContent, f.presenter.Content)
	})

	t.Run("It validates empty component", func(t *testing.T) {
		f := setup()

		payload := &Component{}
		service := NewRenderPayloadService(payload, f.renderer, f.presenter)
		service.Execute()

		assert.False(t, f.renderer.RenderCalled)
		assert.DeepEqual(t, invalidPayloadErrors, f.presenter.ValidationErrors)
	})

	t.Run("It validates nil component", func(t *testing.T) {
		f := setup()

		service := NewRenderPayloadService(nil, f.renderer, f.presenter)
		service.Execute()

		assert.False(t, f.renderer.RenderCalled)
		assert.DeepEqual(t, invalidPayloadErrors, f.presenter.ValidationErrors)
	})
}

type RendererSpy struct {
	Component    *Component
	Content      string
	RenderCalled bool
}

func NewRendererSpy() *RendererSpy {
	return &RendererSpy{}
}

func (r *RendererSpy) Render(c *Component) string {
	r.RenderCalled = true
	r.Component = c
	return r.Content
}

type ContentPresenterSpy struct {
	Content          string
	ValidationErrors []ValidationError
}

func NewContentPresenterSpy() *ContentPresenterSpy {
	return &ContentPresenterSpy{}
}

func (p *ContentPresenterSpy) PresentContent(content string) {
	p.Content = content
}

func (p *ContentPresenterSpy) PresentValidationErrors(errs []ValidationError) {
	p.ValidationErrors = errs
}
