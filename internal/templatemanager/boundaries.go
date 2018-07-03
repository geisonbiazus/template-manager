package templatemanager

type RenderTemplateOutputBoundary interface {
	PresentHTML(html string)
	PresentValidationErrors([]ValidationError)
}

type RenderTemplateInputBoundary interface {
	RenderByJSON(template *Component, output RenderTemplateOutputBoundary)
}

type ManageTemplateInputBoundary interface {
	Create(template Template, output ManageTemplateOutputBoundary)
}

type ManageTemplateOutputBoundary interface {
	PresentCreated(template Template)
	PresentValidationErrors([]ValidationError)
}
