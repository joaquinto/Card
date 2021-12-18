package services

import (
	"card/domain"
	"errors"
)

// OpenDeck This implements the open deck service feature
func (cs *CardService) OpenDeck(deckID string) (*domain.Deck, error) {
	var found bool
	var deck *domain.Deck

	for _, d := range Decks {
		if d.ID == deckID {
			found = true
			deck = d
			break
		}
	}

	if !found {
		return &domain.Deck{}, errors.New("Deck Not Found")
	}
	
	return deck, nil
}
