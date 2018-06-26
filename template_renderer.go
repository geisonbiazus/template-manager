package templatemanager

type TemplateRenderer struct {
}

func NewTemplateRenderer() *TemplateRenderer {
	return &TemplateRenderer{}
}

func (r *TemplateRenderer) Render(component *Component) string {
	return component.Render()
}
