package services

import (
	"card/domain"
	"testing"

	"github.com/stretchr/testify/require"
)

type drawTestCard struct {
	deckId        string
	count         uint8
	selectedCard  []*domain.Card
	errorResponse string
}

func TestDrawCard(t *testing.T) {
	cardService := NewCardService()

	sc := CreateCard()
	ssc := ShuffleCard(CreateCard())
	mcs := MakeCardSelection(CreateCard(), []string{"AS", "KD", "AC", "2C", "KH"})
	smcs := ShuffleCard(MakeCardSelection(CreateCard(), []string{"AS", "KD", "AC", "2C", "KH", "3C"}))

	Decks = []*domain.Deck{
		{ID: "414c59ae-51a9-4d15-a1bc-8983ccdeea19", Shuffled: false, Remaining: 52, Cards: sc},
		{ID: "4f84dabf-f4b6-4282-8ab9-d7a1a5cb7700", Shuffled: true, Remaining: 52, Cards: ssc},
		{ID: "c7f83a96-6a16-46b3-84cb-4e4c09ce6505", Shuffled: false, Remaining: 5, Cards: mcs},
		{ID: "28041656-40f3-46b0-bf91-4580fc7a8ccb", Shuffled: true, Remaining: 6, Cards: smcs},
	}

	testTable := []drawTestCard{
		{deckId: "414c59ae-51a9-4d15-a1bc-8983ccdeea19", count: 2, selectedCard: sc[:2], errorResponse: "Deck Not Found"},
		{deckId: "4f84dabf-f4b6-4282-8ab9-d7a1a5cb7700", count: 5, selectedCard: ssc[:5], errorResponse: "Deck Not Found"},
		{deckId: "c7f83a96-6a16-46b3-84cb-4e4c09ce6505", count: 2, selectedCard: mcs[:2], errorResponse: "Deck Not Found"},
		{deckId: "28041656-40f3-46b0-bf91-4580fc7a8ccb", count: 3, selectedCard: smcs[:3], errorResponse: "Deck Not Found"},
		{deckId: "28041656-40f3-46b0-bf91-4580fc7a8cce", count: 3, selectedCard: nil, errorResponse: "Deck Not Found"},
	}

	for _, tt := range testTable {
		d, err := cardService.DrawCard(tt.deckId, tt.count)
		if err != nil {
			require.Equal(t, tt.errorResponse, err.Error())
		}

		require.Equal(t, tt.selectedCard, d)
		require.Equal(t, len(tt.selectedCard), len(d))
	}
}
