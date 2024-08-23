package middleware

import (
    "net/http"
    "fmt"
)

func Auth(h http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("Hi mom!")
        h.ServeHTTP(w,r)
    })
}
