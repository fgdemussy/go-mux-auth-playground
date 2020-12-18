package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", SomeHandler)
	r.Use(basicAuthMiddleware)
	log.Fatal(http.ListenAndServe(":8000", r))
}

// SomeHandler does this
func SomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("handled OK!\n"))
}

func basicAuthMiddleware(next http.Handler) http.Handler  {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("www-Authenticate", `Basic realm="Restricted"`)

		username, password, _ := r.BasicAuth()

		if username != "fran" || password != "pass" {
			http.Error(w, "Not Authorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}