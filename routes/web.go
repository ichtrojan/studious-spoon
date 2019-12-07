package routes

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

func Init() *mux.Router {
	route := mux.NewRouter()

	route.NotFoundHandler = http.HandlerFunc(notFound)

	return route
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)

	view, err := template.ParseFiles("views/errors/404.html")

	if err != nil {
		log.Fatal(err)
	}

	_ = view.Execute(w, nil)
}
