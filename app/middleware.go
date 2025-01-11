package app

import (
	"net/http"
)

func RootMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html;charset=utf-8")

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
