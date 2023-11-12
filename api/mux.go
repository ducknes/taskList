package api

import (
	"log"
	"net/http"
	"taskList/handlers"
	"taskList/service"

	"github.com/gorilla/mux"
)

func NewRouter(taskService *service.TaskService) http.Handler {
	router := mux.NewRouter()
	router.Use(commonMiddleware)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("hello User")); err != nil {
			log.Println(err)
		}
	}).Methods(http.MethodGet)

	router.HandleFunc("/tasks", handlers.GetAllTasks(taskService)).Methods(http.MethodGet)

	router.HandleFunc("/task/{id}", handlers.GetTask(taskService)).Methods(http.MethodGet)

	router.HandleFunc("/save", handlers.SaveTask(taskService)).Methods(http.MethodPost)

	router.HandleFunc("/delete", handlers.DeleteTask(taskService)).Methods(http.MethodDelete)

	router.HandleFunc("update", handlers.UpdateTask(taskService)).Methods(http.MethodPut)

	return router
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
