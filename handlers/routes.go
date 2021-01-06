package handlers

import (
	"net/http"
)

// RegisterRoutes is use to register routes.
func RegisterRoutes() {
	http.HandleFunc("/default", defaultHandler)
}
