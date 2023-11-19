package main

import (
	"log"
	"taskList/api"
	"taskList/config"
	"taskList/database"
	"taskList/repository"
	"taskList/service"
)

type serviceProvider struct {
	settings   *config.Settings
	doneSignal chan struct{}
}

func (s serviceProvider) provideServer() *api.Server {
	postgres, err := database.NewPostgresConnection(s.settings.PostgresConnection)
	if err != nil {
		log.Fatal(err)
	}

	taskRepository := repository.NewTaskRepository(postgres)

	taskService := service.NewTaskService(taskRepository)

	router := api.NewRouter(taskService)

	return api.NewTaskListServer(s.settings, router, s.doneSignal)
}
