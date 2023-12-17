package handlers

import (
	"encoding/json"
	"net/http"
	"taskList/service"
	"taskList/service/models"
)

func GetAllTasks(taskService *service.TaskService) http.HandlerFunc {
	return apiHandler(func(r *http.Request) (interface{}, error) {
		return taskService.GetAllTasks()
	})
}

func GetTask(taskService *service.TaskService) http.HandlerFunc {
	return apiHandler(func(r *http.Request) (interface{}, error) {
		id, err := parseVarsId(r)
		if err != nil {
			return nil, err
		}

		return taskService.GetTask(id)
	})
}

func SaveTask(taskService *service.TaskService) http.HandlerFunc {
	return apiHandler(func(r *http.Request) (interface{}, error) {
		var newTask models.Task
		if err := json.NewDecoder(r.Body).Decode(&newTask); err != nil {
			return nil, err
		}

		return taskService.AddTask(newTask)
	})

}

func DeleteTask(taskService *service.TaskService) http.HandlerFunc {
	return apiHandler(func(r *http.Request) (interface{}, error) {
		id, err := parseVarsId(r)
		if err != nil {
			return nil, err
		}

		return taskService.DeleteTask(id)
	})
}

func UpdateTask(taskService *service.TaskService) http.HandlerFunc {
	return apiHandler(func(r *http.Request) (interface{}, error) {
		var updatedTask models.Task
		if err := json.NewDecoder(r.Body).Decode(&updatedTask); err != nil {
			return nil, err
		}

		return taskService.UpdateTask(updatedTask)
	})
}
