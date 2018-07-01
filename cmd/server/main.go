package main

import (
	"net/http"

	"github.com/geisonbiazus/templatemanager"
	"github.com/geisonbiazus/templatemanager/cmd/server/handlers"
	"github.com/geisonbiazus/templatemanager/cmd/server/presenters"
)

func main() {
	mux := http.NewServeMux()

	templateRenderer := templatemanager.NewTemplateRenderer("test/templates/*")
	renderTemplateInteractor := templatemanager.NewRenderTemplateInteractor(templateRenderer)
	renderTemplateJSONPresenterFactory := presenters.NewRenderTemplateJSONPresenterFactory()

	renderTemplateByJSONHandler := handlers.NewRenderTemplateByJSONHandler(
		renderTemplateInteractor, renderTemplateJSONPresenterFactory,
	)

	mux.Handle("/render_by_json", renderTemplateByJSONHandler)

	http.ListenAndServe(":3001", mux)
}
