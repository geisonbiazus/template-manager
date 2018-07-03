package templatemanager

type Template struct {
	Body *Component `json:"body"`
}

type Properties map[string]string

type Component struct {
	Type       string       `json:"type"`
	Children   []*Component `json:"children"`
	Properties Properties   `json:"properties"`
}

func (c *Component) Empty() bool {
	return c.Type == "" && len(c.Children) == 0
}

const (
	ErrorInvalid = "INVALID"
)

type ValidationError struct {
	Field   string `json:"field"`
	Type    string `json:"type"`
	Message string `json:"message"`
}
