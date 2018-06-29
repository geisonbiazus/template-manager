package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/geisonbiazus/templatemanager"
)

func main() {
	template := &templatemanager.Component{}
	json.NewDecoder(os.Stdin).Decode(template)

	presenter := NewFilePresenter(os.Stdout)

	renderer := templatemanager.NewTemplateRenderer("test/templates/*")
	interactor := templatemanager.NewRenderTemplateInteractor(renderer)
	interactor.RenderByJSON(template, presenter)
}

type FilePresenter struct {
	Writer io.Writer
}

func NewFilePresenter(w io.Writer) *FilePresenter {
	return &FilePresenter{
		Writer: w,
	}
}

func (p *FilePresenter) PresentHTML(html string) {
	fmt.Fprintln(p.Writer, html)
}

func (p *FilePresenter) PresentValidationErrors(errors []templatemanager.ValidationError) {
	for _, error := range errors {
		fmt.Fprintln(p.Writer, error.Message)
	}
}
