package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {

	router := chi.NewRouter()

	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	router.Get("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("User List"))
	})
}
