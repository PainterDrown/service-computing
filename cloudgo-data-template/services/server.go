package services

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
	n := negroni.Classic()
	mx := mux.NewRouter()
	initRoutes(mx, formatter)
	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/tests/{any}", testHandler(formatter)).Methods("GET")
	mx.HandleFunc("/users", postUserInfoHandler(formatter)).Methods("POST")
	mx.HandleFunc("/users", getUserInfoHandler(formatter)).Methods("GET")
}

func testHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		any := vars["any"]
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"Hello " + any})
	}
}
