package main

import (
	"context"
	"encoding/json"
	"flag"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/griggsca91/quick-go-demo/repository"
	"github.com/griggsca91/quick-go-demo/repository/cards"
)

func main() {

	repo := flag.String("repo", "fake", "choose which repo to use")

	flag.Parse()

	var cardsRepo cards.CardRepository

	if *repo != "fake" {
		databaseURL := os.Getenv("DATABASE_URL")

		var err error
		cardsRepo, err = repository.NewPostgresRepository(databaseURL)
		if err != nil {
			panic(err)
		}
	} else {
		cardsRepo = repository.NewFakeRepository()
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
