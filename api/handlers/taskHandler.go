package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"taskList/service"
)

func GetAllTasks(taskService *service.TaskService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func GetTask(taskService *service.TaskService) http.HandlerFunc {
	return apiHandler(func(r *http.Request) (interface{}, error) {
		vars := mux.Vars(r)
		id, ok := vars["id"]
		if !ok {
			return nil, fmt.Errorf("не удалось получить id")
		}

		return taskService.GetTask(id)
	})
}

func SaveTask(taskService *service.TaskService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func DeleteTask(taskService *service.TaskService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func UpdateTask(taskService *service.TaskService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
