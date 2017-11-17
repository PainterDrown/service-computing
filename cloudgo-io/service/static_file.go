package service

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// StaticFile listens "/assets" for static file.
func StaticFile(m *mux.Router) {
	webRoot := os.Getenv("WEBROOT")
	if len(webRoot) == 0 {
		if root, err := os.Getwd(); err != nil {
			panic("Could not retrive working directory")
		} else {
			webRoot = root
		}
	}
	m.PathPrefix("/assets").Handler(http.FileServer(http.Dir(webRoot + "/")))
}
