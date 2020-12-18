package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", SomeHandler)
	log.Fatal(http.ListenAndServe(":8000", r))
}

// SomeHandler does this
func SomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Gorilla!\n"))
}