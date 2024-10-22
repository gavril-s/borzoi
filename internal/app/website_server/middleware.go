package websiteserver

import (
	"fmt"
	"net/http"
	"time"
)

func (s *server) middleware(next http.Handler, timeout time.Duration) http.Handler {
	return http.TimeoutHandler(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Date", time.Now().UTC().Format(http.TimeFormat))
			w.Header().Set("X-Request-Time", time.Now().Format(time.RFC1123))
			w.Header().Set("Content-Type", "text/html")
			if r.Method != http.MethodGet {
				w.WriteHeader(http.StatusMethodNotAllowed)
				fmt.Fprintf(w, "Method not allowed")
			}
			next.ServeHTTP(w, r)
		}),
		timeout,
		"Request timed out",
	)
}
