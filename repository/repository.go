package repository

import (
	"context"
	"fmt"

	"github.com/go-errors/errors"
	"github.com/griggsca91/quick-go-demo/repository/cards"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresCard struct {
	CardName string
}

func (p PostgresCard) Name() string {
	return p.CardName
}

type PostgresRepository struct {
	pool *pgxpool.Pool
}

func (p PostgresRepository) GetCards(ctx context.Context, name string) ([]cards.Card, error) {

	rows, err := p.pool.Query(ctx, "select name from cards where name like $1 || '%'", name)
	if err != nil {
		fmt.Println("transaction error", err.Error())
		return nil, err
	}

	var mtgCards []cards.Card
	for rows.Next() {
		var name string
		rows.Scan(&name)
		mtgCards = append(mtgCards, PostgresCard{
			CardName: name,
		})
	}

	return mtgCards, nil
}

func NewPostgresRepository(databaseURL string) (PostgresRepository, error) {
	pool, err := pgxpool.New(context.Background(), databaseURL)
	if err != nil {
		return PostgresRepository{}, errors.New(err)
	}

	return PostgresRepository{
		pool: pool,
	}, pool.Ping(context.Background())
}
