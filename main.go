package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {

	fmt.Println("Hello world")

	router := chi.NewRouter()

	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			queryParams := r.URL.Query()
			if !queryParams.Has("password") {
				http.Error(w, "plz no hacking", http.StatusForbidden)
				return
			}
			next.ServeHTTP(w, r)
		})
	})

	router.Get("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		healthcheckResponse := struct {
			Message string `json:"message"`
		}{
			Message: "Hello World",
		}

		bytes, err := json.Marshal(healthcheckResponse)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(bytes)
	})

	panic(http.ListenAndServe(":3333", router))
}
