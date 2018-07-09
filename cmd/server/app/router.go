package app

import (
	"net/http"

	"github.com/geisonbiazus/templatemanager/cmd/server/handler"
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

	mux.Handle("/v1/templates/render", renderTemplateByJSONHandler)

	return mux
}
