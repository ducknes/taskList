package app

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"taskList/api"
)

func Start() {
	server := api.NewServer()

	go func() {
		if err := server.ListenAndServe(); err != nil {
			panic(err)
		}
	}()
	log.Println(fmt.Sprintf("server is running on addr: %s", server.Addr))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	log.Println("server is stopped")
}
