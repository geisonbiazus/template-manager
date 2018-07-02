package app

import (
	"net/http"

	"github.com/geisonbiazus/templatemanager/cmd/server/handlers"
	"github.com/geisonbiazus/templatemanager/cmd/server/presenters"
	"github.com/geisonbiazus/templatemanager/internal/templatemanager"
)

func Mux(templatePath string) http.Handler {
	mux := http.NewServeMux()

	templateRenderer := templatemanager.NewTemplateRenderer(templatePath)
	renderTemplateInteractor := templatemanager.NewRenderTemplateInteractor(templateRenderer)
	renderTemplateJSONPresenterFactory := presenters.NewRenderTemplateJSONPresenterFactory()

	renderTemplateByJSONHandler := handlers.NewRenderTemplateByJSONHandler(
		renderTemplateInteractor, renderTemplateJSONPresenterFactory,
	)

	mux.Handle("/render_by_json", renderTemplateByJSONHandler)

	return mux
}
