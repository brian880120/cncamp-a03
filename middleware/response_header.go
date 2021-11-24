package middleware

import (
	"net/http"
	"os"
)

func ResponseHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		for name, headers := range req.Header {
			for _, h := range headers {
				w.Header().Add(name, h)
			}
		}

		version := os.Getenv("VERSION")
		w.Header().Add("version", version)
		next.ServeHTTP(w, req)
	})
}
