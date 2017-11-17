package service

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// JSONRequest listens "/api" for JSON request.
func JSONRequest(m *mux.Router) {
	r := render.New(render.Options{
		IndentJSON: true,
	})
	m.HandleFunc("/json", func(w http.ResponseWriter, req *http.Request) {
		r.JSON(w, http.StatusOK, map[string]string{
			"author": "painterdrown",
			"email":  "painterdrown@hotmail.com",
			"blog":   "https://painterdrown.github.io",
			"github": "https://github.com/painterdrown",
		})
	})
}
