package main

import (
	"net/http"

	"github.com/geisonbiazus/templatemanager"
)

func main() {
	mux := http.NewServeMux()

	templateRenderer := templatemanager.NewTemplateRenderer("test/templates/*")
	renderTemplateInteractor := templatemanager.NewRenderTemplateInteractor(templateRenderer)
	renderTemplateJSONPresenterFactory := templatemanager.NewRenderTemplateJSONPresenterFactory()

	renderTemplateByJSONHandler := templatemanager.NewRenderTemplateByJSONHandler(
		renderTemplateInteractor, renderTemplateJSONPresenterFactory,
	)

	mux.Handle("/render_by_json", renderTemplateByJSONHandler)

	http.ListenAndServe(":3001", mux)
}
