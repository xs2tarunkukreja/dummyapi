package helper

import (
	"net/http"
)

// MyHandler is a function type.
type MyHandler func(http.ResponseWriter, *http.Request)

var routes = make(map[string]MyHandler)

// RegisterRoutes is use to register routes.
func RegisterRoutes(url string, handler MyHandler) {
	routes[url] = handler
}

// Routes register routes with http package.
func Routes() {
	for url, handler := range routes {
		http.HandleFunc(url, handler)
	}
}
