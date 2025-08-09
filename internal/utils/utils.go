package utils

import (
	"bytes"
	"html/template"
	"net/http"
	"tb/internal/app"
	"tb/internal/models"
)

func InternalServerError(w http.ResponseWriter) {
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}

func SendHtmlResponse(w http.ResponseWriter, a *app.App, data models.Page) {
	t, err := a.LoadTemplate("index.html")
	if err != nil {
		InternalServerError(w)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	err = t.Execute(w, data)
	if err != nil {
		InternalServerError(w)
		return
	}
}

func RenderTemplateToHTML(t *template.Template, data any) (template.HTML, error) {
	var buf bytes.Buffer
	err := t.Execute(&buf, data)
	if err != nil {
		return "", err
	}
	return template.HTML(buf.String()), nil
}
