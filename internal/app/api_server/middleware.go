package apiserver

import (
	"fmt"
	"net/http"
	"time"
)

func (s *server) middleware(next http.Handler, timeout time.Duration, method string) http.Handler {
	return http.TimeoutHandler(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Date", time.Now().UTC().Format(http.TimeFormat))
			w.Header().Set("X-Request-Time", time.Now().Format(time.RFC1123))
			w.Header().Set("Content-Type", "application/json")
			if r.Method != method {
				w.WriteHeader(http.StatusMethodNotAllowed)
				fmt.Fprintf(w, "Method not allowed")
				return
			}
			next.ServeHTTP(w, r)
		}),
		timeout,
		"Request timed out",
	)
}

func (s *server) postMiddleware(next http.Handler, timeout time.Duration) http.Handler {
	return s.middleware(next, timeout, http.MethodPost)
}

func (s *server) getMiddleware(next http.Handler, timeout time.Duration) http.Handler {
	return s.middleware(next, timeout, http.MethodGet)
}
