package middleware

import "net/http"

func RateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: check rate limit in Valkey
		next.ServeHTTP(w, r)
	})
}
