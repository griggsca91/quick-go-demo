package cards

import "context"

type Card interface {
	Name() string
}

type CardRepository interface {
	GetCards(ctx context.Context, name string) ([]Card, error)
}
