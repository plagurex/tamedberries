package handlers

import (
	"net/http"
	"tb/internal/app"
)

func HomeHandler(a *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, web!"))
	}
}
