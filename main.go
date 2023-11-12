package main

import (
	"log"
	"taskList/config"
)

func main() {
	log.SetFlags(log.Flags() | log.Llongfile)
	settings := config.Read()
	config.Init()

	doneSignal := make(chan struct{})

	server := serviceProvider{
		settings:   settings,
		doneSignal: doneSignal,
	}.provideServer()

	server.Start()

	<-doneSignal
}
