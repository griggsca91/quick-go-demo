package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/griggsca91/quick-go-demo/repository"
)

func main() {

	fmt.Println("Hello world")

	databaseURL := os.Getenv("DATABASE_URL")

	cardsRepo, err := repository.NewPostgresRepository(databaseURL)
	if err != nil {
		panic(err)
	}

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

	router.Get("/cards", func(w http.ResponseWriter, r *http.Request) {
		queryParams := r.URL.Query()
		if !queryParams.Has("cardName") {
			http.Error(w, "plz come back with a name and stop wasting my time", http.StatusForbidden)
			return
		}

		cardName := queryParams.Get("cardName")

		cards, err := cardsRepo.GetCards(context.Background(), cardName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		type Card struct {
			Name string `json:"name"`
		}

		type cardsResponse struct {
			Cards []Card `json:"cards"`
		}

		var response cardsResponse

		for _, card := range cards {
			response.Cards = append(response.Cards, Card{
				Name: card.Name(),
			})
		}

		bytes, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(bytes)

	})

	panic(http.ListenAndServe(":3333", router))
}
