package utils

import "net/http"

// JsonContentTypeMiddleware sets the Content-Type header to "application/json"
// for all responses and then calls the next handler in the chain.
func JsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
