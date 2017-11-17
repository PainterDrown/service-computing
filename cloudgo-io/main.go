package main

import (
	"service-computing/cloudgo-io/service"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func main() {
	n := negroni.Classic()
	m := mux.NewRouter()

	// listens "/assets" for static file
	service.StaticFile(m)
	// listens "/json" for JSON request
	service.JSONRequest(m)
	// listens "/form" for form request
	service.FormRequest(m)
	// listens "/unknown" for not implemented
	service.NotImplemented(m)

	n.UseHandler(m)
	n.Run(":2222")
}
