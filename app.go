package main

import (
	"taskList/api"
	"taskList/config"
)

type serviceProvider struct {
	settings   *config.Settings
	doneSignal chan struct{}
}

func (s serviceProvider) provideServer() *api.Server {
	//TODO: repos and db init

	router := api.NewRouter()

	return api.NewTaskListServer(s.settings, router, s.doneSignal)
}
