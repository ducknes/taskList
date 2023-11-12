package api

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"taskList/config"
)

func NewTaskListServer(settings *config.Settings, mux http.Handler, done chan struct{}) *Server {
	return &Server{
		done:     done,
		settings: settings,
		mux:      mux,
	}
}

type Server struct {
	done       chan struct{}
	settings   *config.Settings
	mux        http.Handler
	httpServer *http.Server
}

func (s *Server) Start() {
	host := fmt.Sprintf(":%s", s.settings.Port)
	listener, err := net.Listen("tcp", host)
	if err != nil {
		log.Println(err)
		panic("TCP listener not created")
	}

	s.httpServer = &http.Server{
		Addr:    host,
		Handler: s.mux,
	}
	go s.httpServer.Serve(listener)
	log.Printf("server started on %s", s.settings.Port)
}

func (s *Server) Stop() {
	err := s.httpServer.Close()
	if err != nil {
		log.Println(err)
		panic("can`t clone server")
	}

	log.Println("server stopped")
	if s.done != nil {
		s.done <- struct{}{}
	}
}
