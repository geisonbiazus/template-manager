package app

import (
	"net/http"

	"github.com/geisonbiazus/templatemanager/cmd/server/handlers/template"
	"github.com/geisonbiazus/templatemanager/internal/templatemanager"
	"github.com/geisonbiazus/templatemanager/internal/templatemanager/rendertemplate"
)

func Mux(templatePath string) http.Handler {
	mux := http.NewServeMux()

	templateRenderer := templatemanager.NewTemplateRenderer(templatePath)
	renderTemplateInteractor := rendertemplate.NewInteractor(templateRenderer)
	renderByJSONHandler := template.NewRenderByJSONHandler(renderTemplateInteractor)

	mux.Handle("/v1/templates/render", renderByJSONHandler)

	return mux
}
