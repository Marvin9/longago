package middleware

import (
	"net/http"
)

// SetJSONHeaderMiddleware set header content type for json
func SetJSONHeaderMiddleware(fun http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("content-type", "application/json")
		fun(w, req)
	}
}
