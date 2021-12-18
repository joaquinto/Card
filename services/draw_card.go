package services

import (
	"card/domain"
	"errors"
)

// DrawCard This implements the draw card service feature
func (cs *CardService) DrawCard(deckID string, noOfCard uint8) ([]*domain.Card, error) {
	var drawnCards []*domain.Card
	var found bool

	if noOfCard < 1 {
		return drawnCards, errors.New("Count must be equals to or greater than 1")
	}

	for _, d := range Decks {
		if d.ID == deckID {
			if d.Remaining == 0 {
				return drawnCards, errors.New("There are no card left")
			}
			found = true
			drawnCards = drawCard(d, noOfCard)
			break
		}
	}

	if !found {
		return drawnCards, errors.New("Deck Not Found")
	}

	return drawnCards, nil
}

// drawCard This implements the logic to picking the first number of 
// card as specified by the used and passed in the query parameter as count
func drawCard(deck *domain.Deck, noOfCard uint8) []*domain.Card {
	var i uint8
	var cards []*domain.Card
	for i = 0; i < noOfCard; i++ {
		cards = append(cards, deck.Cards[0])
		deck.Cards = deck.Cards[1:]
		deck.Remaining -= 1
	}

	return cards
}
