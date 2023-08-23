package repository

import (
	"context"
	"fmt"

	"github.com/griggsca91/quick-go-demo/repository/cards"
)

type FakeCard struct {
	fakeName string
}

func (f FakeCard) Name() string {
	return fmt.Sprintf("Fake name - %s", f.fakeName)
}

type FakeRepository struct{}

var _ cards.CardRepository = FakeRepository{}

// If you create a new function, you don't really need the line above

func NewFakeRepository() cards.CardRepository {
	return FakeRepository{}
}

func (FakeRepository) GetCards(_ context.Context, cardName string) ([]cards.Card, error) {

	var response []cards.Card

	for i := 0; i < 10; i++ {
		response = append(response, FakeCard{
			fakeName: fmt.Sprintf("Fake Card %v", i),
		})
	}

	return response, nil
}
