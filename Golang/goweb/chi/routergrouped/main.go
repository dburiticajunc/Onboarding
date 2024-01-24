package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	router := chi.NewRouter()

	// with subrouter
	router.Route("/users", func(r chi.Router) {
		// users
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("users"))

		})
		// user by id
		r.Get("/{userId}", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("user by id"))
		})
		// posts by user id
		r.Get("/{userId}/posts", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("posts by user id"))
		})

	})

	http.ListenAndServe(":8080", router)
}
