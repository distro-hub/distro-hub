package middleware

import (
	"fmt"
	"net/http"
)

func Auth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hi mom!")
		h.ServeHTTP(w, r)
	})
}
