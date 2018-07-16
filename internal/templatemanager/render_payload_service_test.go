package templatemanager

import (
	"testing"

	"github.com/geisonbiazus/templatemanager/internal/support/assert"
)

func TestRenderPayloadInteractor(t *testing.T) {
	t.Run("It renders a payload and presents it", func(t *testing.T) {
		component := &Component{Type: "Page"}

		renderer := NewRendererSpy()
		presenter := NewContentPresenterSpy()
		interactor := NewRenderPayloadService(component, renderer, presenter)

		expectedContent := "RenderedContent"
		renderer.Content = expectedContent

		interactor.Execute()

		assert.DeepEqual(t, component, renderer.Component)
		assert.Equal(t, expectedContent, presenter.Content)
	})
}

type RendererSpy struct {
	Component *Component
	Content   string
}

func NewRendererSpy() *RendererSpy {
	return &RendererSpy{}
}

func (r *RendererSpy) Render(c *Component) string {
	r.Component = c
	return r.Content
}

type ContentPresenterSpy struct {
	Content string
}

func NewContentPresenterSpy() *ContentPresenterSpy {
	return &ContentPresenterSpy{}
}

func (p *ContentPresenterSpy) PresentContent(content string) {
	p.Content = content
}
