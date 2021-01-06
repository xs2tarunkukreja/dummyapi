package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/xs2tarunkukreja/dummyapi/helper"
)

func init() {
	helper.RegisterRoutes("/default", defaultHandler)
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	helloMessage := struct {
		Message string
	}{
		Message: "Hello World V 1.0",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(&helloMessage)
}
