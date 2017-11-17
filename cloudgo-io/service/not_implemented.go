package service

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// NotImplemented listens "/unknown" for not implemented
func NotImplemented(m *mux.Router) {
	r := render.New()
	m.HandleFunc("/unknown", func(w http.ResponseWriter, req *http.Request) {
		r.HTML(w, http.StatusNotImplemented, "error", "Not Implemented")
	})
}
