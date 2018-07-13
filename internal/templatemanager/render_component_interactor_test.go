package templatemanager

import "testing"

func TestRenderPayloadInteractor(t *testing.T) {
	t.Run("Render a component and present it", func(t *testing.T) {
		input := RenderPayloadInput{
			Component: &Component{Type: "Page"},
		}

		presenter := NewContentPresenterSpy()
		interactor := NewRenderPayloadInteractor(input, presenter)
		interactor.Execute()
	})
}

type ContentPresenterSpy struct {
}

func NewContentPresenterSpy() *ContentPresenterSpy {
	return &ContentPresenterSpy{}
}
