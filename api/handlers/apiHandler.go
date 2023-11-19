package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func apiHandler(action func(request *http.Request) (interface{}, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
