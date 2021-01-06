package handlers

import (
	"encoding/json"
	"net/http"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	helloMessage := struct {
		Message string
	} {
		Message: "Hello World",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(&helloMessage)
}
