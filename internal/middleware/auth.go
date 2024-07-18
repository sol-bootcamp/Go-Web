package middleware

import (
	"net/http"
)

func Auth(apiToken string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Otra forma de obtener el token
			// apiToken :=os.Getenv("API_TOKEN")
			token := r.Header.Get("Authorization")
			if token != apiToken {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized"))
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
