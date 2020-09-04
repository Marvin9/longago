package middleware

import (
	"net/http"
)

func SetJSONHeaderMiddleware(fun http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("content-type", "application/json")
		fun(w, req)
	}
}
