package main

import (
	"net/http"

	"github.com/geisonbiazus/templatemanager/cmd/server/app"
	"github.com/geisonbiazus/templatemanager/internal/templatemanager"
)

func main() {
	mux := app.Mux(templatemanager.DefaultTemplatePath)
	http.ListenAndServe(":3001", mux)
}
