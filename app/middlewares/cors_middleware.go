package middlewares

import (
	"github.com/leandrose/uptime-kuma-api-go/config"
	"net/http"
	"strings"
)

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c := config.GetConfig()
		if len(c.CrossOrigin) > 0 {
			w.Header().Set("Access-Control-Allow-Origin", strings.Join(c.CrossOrigin, ","))
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")

			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
			}
		}

		next.ServeHTTP(w, r)
	})
}
