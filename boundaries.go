package templatemanager

type RenderTemplateOutputBoundary interface {
	PresentHTML(html string)
	PresentValidationErrors([]ValidationError)
}

type RenderTemplateInputBoundary interface {
	RenderByJSON(template *Component, presenter RenderTemplateOutputBoundary)
}
