package api

import "net/http"

func NewServer() http.Server {
	router := newRouter()

	return http.Server{
		Addr:    ":8080",
		Handler: router,
	}
}
