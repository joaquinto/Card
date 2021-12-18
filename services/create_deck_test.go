package services

import (
	"card/domain"
	"testing"

	"github.com/stretchr/testify/require"
)

type createDeck struct {
	shuffle   bool
	selection []string
	cardSize  uint8
	cards     []*domain.Card
}

func TestCreateDeck(t *testing.T) {
	cardService := NewCardService()
	cards := CreateCard()

	testTable := []createDeck{
		{shuffle: false, selection: []string{}, cardSize: 52, cards: cards},
		{shuffle: true, selection: []string{}, cardSize: 52, cards: cards},
		{shuffle: false, selection: []string{"AS","KD","AC","2C","KH"}, cardSize: 5, cards: MakeCardSelection(cards, []string{"AS","KD","AC","2C","KH"})},
		{shuffle: true, selection: []string{"AS","KD","AC","2C","KH", "3C"}, cardSize: 6, cards: MakeCardSelection(cards, []string{"AS","KD","AC","2C","KH", "3C"})},
	}

	for _, tt := range testTable {
		deck := cardService.CreateDeck(tt.shuffle, tt.selection)
		require.Equal(t, tt.cardSize, deck.Remaining)
		if tt.shuffle {
			require.NotEqual(t, tt.cards, deck.Cards)
		} else {
			require.Equal(t, tt.cards, deck.Cards)
		}
	}
}
