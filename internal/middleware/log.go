package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("Request URI: ", r.RequestURI+" "+r.Method, "Resquest recieved at: ", time.Now().Format("2006-01-02 15:04:05"))

		next.ServeHTTP(w, r)
	})
}
