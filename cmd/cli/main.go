package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/geisonbiazus/templatemanager"
)

func main() {
	content, _ := ioutil.ReadAll(os.Stdin)
	presenter := NewFilePresenter(os.Stdout)

	renderer := templatemanager.NewTemplateRenderer("test/templates/*")
	interactor := templatemanager.NewRenderTemplateInteractor(renderer)
	interactor.RenderByJSON(string(content), presenter)
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
