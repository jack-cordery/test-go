package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type StatusResponse struct {
	OldMessage string
	NewMessage string
}

func getStatus() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(StatusResponse{
			OldMessage: "Who?",
			NewMessage: "When?",
		})
		if err != nil {
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
			return
		}
	}
}


func main() {
	 mux := http.NewServeMux()

	 mux.HandleFunc("GET /status", getStatus())

	 err := http.ListenAndServe(":8001", mux)
	if err != nil {
		log.Fatal(err)
	}

}