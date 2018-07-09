package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/geisonbiazus/templatemanager/internal/templatemanager"
	"github.com/geisonbiazus/templatemanager/internal/templatemanager/rendertemplate"
)

func main() {
	template := &templatemanager.Component{}
	json.NewDecoder(os.Stdin).Decode(template)

	renderer := templatemanager.NewTemplateRenderer("internal/templatemanager/test/templates/*")
	interactor := rendertemplate.NewInteractor(renderer)
	resp := interactor.RenderByJSON(
		rendertemplate.RenderByJSONRequest{rendertemplate.Template{Body: template}},
	)

	if resp.Status == rendertemplate.StatusSuccess {
		fmt.Fprintln(os.Stdout, resp.HTML)
	} else if resp.Status == rendertemplate.StatusInvalid {
		for _, error := range resp.Errors {
			fmt.Fprintln(os.Stdout, error.Message)
		}
	}
}
