package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func newRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("hello User")); err != nil {
			log.Println(err)
		}
	}).Methods(http.MethodGet)

	router.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {}).Methods(http.MethodGet)

	router.HandleFunc("/task/{id}", func(w http.ResponseWriter, r *http.Request) {}).Methods(http.MethodGet)

	router.HandleFunc("/save", func(w http.ResponseWriter, r *http.Request) {}).Methods(http.MethodPost)

	return router
}
