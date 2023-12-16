package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func apiHandler(action func(request *http.Request) (interface{}, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Content-Type", "application/json")

		response, err := action(r)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			log.Printf("Не удалось записать ответ: %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}

func parseVarsId(r *http.Request) (string, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return id, fmt.Errorf("не удалось получить id")
	}

	return id, nil
}
