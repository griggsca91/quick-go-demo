package repository

import (
	"context"
	"reflect"
	"testing"

	"github.com/griggsca91/quick-go-demo/repository/cards"
)

func TestFakeRepository_GetCards(t *testing.T) {

	expectedCards := []cards.Card{
		FakeCard{fakeName: "Fake Card 0"},
		FakeCard{fakeName: "Fake Card 1"},
	}

	type args struct {
		in0      context.Context
		cardName string
	}
	tests := []struct {
		name    string
		args    args
		want    []cards.Card
		wantErr bool
	}{
		{
			"always return two cards",
			args{nil, "doesn't matter"},
			expectedCards,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := FakeRepository{}
			got, err := f.GetCards(tt.args.in0, tt.args.cardName)
			if (err != nil) != tt.wantErr {
				t.Errorf("FakeRepository.GetCards() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FakeRepository.GetCards() = %v, want %v", got, tt.want)
			}
		})
	}
}
