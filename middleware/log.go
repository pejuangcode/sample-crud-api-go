package middleware

import (
	"fmt"
	"net/http"
)

func Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Ini dari middleware Log....\n")
		fmt.Println(r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
