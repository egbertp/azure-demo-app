package middleware

import (
	"net/http"
)

func CorsMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		w.Header().Set("Strict-Transport-Security", "max-age=63072000; includeSubDomains")

		if r.Method == "OPTIONS" {
			return // Preflight sets headers and we're done
		}
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func ContentTypeJsonMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func CombinedMiddleware(next http.Handler) http.Handler {
	return ContentTypeJsonMiddleware(CorsMiddleware(next))
}
