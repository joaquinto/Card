package services

import (
	"card/domain"
	"testing"

	"github.com/stretchr/testify/require"
)

type openDeck struct {
	deckId   string
	cardSize uint8
	shuffled bool
}

func TestOpenDeck(t *testing.T) {
	cardService := NewCardService()

	Decks = []*domain.Deck{
		{ID: "414c59ae-51a9-4d15-a1bc-8983ccdeea19", Shuffled: false, Remaining: 52, Cards: CreateCard()},
		{ID: "4f84dabf-f4b6-4282-8ab9-d7a1a5cb7700", Shuffled: true, Remaining: 52, Cards: ShuffleCard(CreateCard())},
		{ID: "c7f83a96-6a16-46b3-84cb-4e4c09ce6505", Shuffled: false, Remaining: 5, Cards: MakeCardSelection(CreateCard(), []string{"AS", "KD", "AC", "2C", "KH"})},
		{ID: "28041656-40f3-46b0-bf91-4580fc7a8ccb", Shuffled: true, Remaining: 6, Cards: ShuffleCard(MakeCardSelection(CreateCard(), []string{"AS", "KD", "AC", "2C", "KH", "3C"}))},
		{ID: "62af6132-4f9b-4e13-b004-2a6f0d8573cd", Shuffled: true, Remaining: 5, Cards: ShuffleCard(MakeCardSelection(CreateCard(), []string{"AS", "KD", "AC", "2C", "KH"}))},
		{ID: "61e7bab6-0588-474d-ac3e-4b587f10f25d", Shuffled: false, Remaining: 6, Cards: MakeCardSelection(CreateCard(), []string{"AS", "KD", "AC", "2C", "KH", "3C"})},
		{ID: "ee9ade4b-0579-4a07-b9d9-cec720e5efd5", Shuffled: false, Remaining: 7, Cards: MakeCardSelection(CreateCard(), []string{"AS", "KD", "AC", "2C", "KH", "3C", "4H"})},
		{ID: "6ec0ce07-6b14-4643-a0d6-114a47025078", Shuffled: true, Remaining: 7, Cards: ShuffleCard(MakeCardSelection(CreateCard(), []string{"AS", "KD", "AC", "2C", "KH", "3C", "4H"}))},
		{ID: "2b67df25-94ca-438f-9f01-fec44035ff85", Shuffled: false, Remaining: 9, Cards: MakeCardSelection(CreateCard(), []string{"AS", "KD", "AC", "2C", "KH", "3C", "4H", "JC", "5D"})},
		{ID: "5232c187-3381-4980-b05b-384b44e2a381", Shuffled: true, Remaining: 9, Cards: ShuffleCard(MakeCardSelection(CreateCard(), []string{"AS", "KD", "AC", "2C", "KH", "3C", "4H", "JC", "5D"}))},
	}

	testTable := []openDeck{
		{deckId: "414c59ae-51a9-4d15-a1bc-8983ccdeea19", shuffled: false, cardSize: 52},
		{deckId: "4f84dabf-f4b6-4282-8ab9-d7a1a5cb7700", shuffled: true, cardSize: 52},
		{deckId: "c7f83a96-6a16-46b3-84cb-4e4c09ce6505", shuffled: false, cardSize: 5},
		{deckId: "28041656-40f3-46b0-bf91-4580fc7a8ccb", shuffled: true, cardSize: 6},
		{deckId: "62af6132-4f9b-4e13-b004-2a6f0d8573cd", shuffled: true, cardSize: 5},
		{deckId: "61e7bab6-0588-474d-ac3e-4b587f10f25d", shuffled: false, cardSize: 6},
		{deckId: "ee9ade4b-0579-4a07-b9d9-cec720e5efd5", shuffled: false, cardSize: 7},
		{deckId: "6ec0ce07-6b14-4643-a0d6-114a47025078", shuffled: true, cardSize: 7},
		{deckId: "2b67df25-94ca-438f-9f01-fec44035ff85", shuffled: false, cardSize: 9},
		{deckId: "5232c187-3381-4980-b05b-384b44e2a381", shuffled: true, cardSize: 9},
	} 

	for _, tt := range testTable {
		d, err := cardService.OpenDeck(tt.deckId)
		if err != nil {
			require.Equal(t, "Deck Not Found", err.Error())
		}

		if tt.deckId != d.ID {
			require.NotEmpty(t, tt.deckId, d.ID)
			require.NotEqual(t, tt.cardSize, d.Remaining)
		} else {
			require.Equal(t, tt.deckId, d.ID)
			require.Equal(t, tt.cardSize, d.Remaining)
		}

	}

}
