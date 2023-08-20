package repository

import (
	"context"

	"github.com/griggsca91/quick-go-demo/repository/cards"
)

type PostgresRepository struct{}

type PostgresCard struct {
}

func (PostgresCard) Name() string {
	return "PostgresCard"
}

func (PostgresRepository) GetCards(_ context.Context, name string) ([]cards.Card, error) {
	return []cards.Card{
		PostgresCard{},
		PostgresCard{},
		PostgresCard{},
		PostgresCard{},
	}, nil
}

func NewPostgresRepository() PostgresRepository {
	return PostgresRepository{}
}
