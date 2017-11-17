package service

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// FormRequest listens "/form" for form request
func FormRequest(m *mux.Router) {
	m.HandleFunc("/form", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		if len(r.Form["first_name"]) == 0 {
			fmt.Fprintf(w, "Error: need parameter \"first_name\"")
			return
		}
		if len(r.Form["last_name"]) == 0 {
			fmt.Fprintf(w, "Error: need parameter \"last_name\"")
			return
		}
		fmt.Fprintf(w, "Hello, %s %s!", r.Form["first_name"][0], r.Form["last_name"][0])
	})
}
