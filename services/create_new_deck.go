package services

import (
	"card/domain"
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

var (
	suits  []string = []string{"SPADES", "DIAMONDS", "CLUBS", "HEARTS"}
	values []string = []string{"ACE", "2", "3", "4", "5", "6", "7", "8", "9", "10", "JACK", "QUEEN", "KING"}
	deck   *domain.Deck
	Decks []*domain.Deck
)

// CreateDeck creates a deck of card based on different parameters 
// such as shuffled and the card selection
func (cs *CardService) CreateDeck(shuffle bool, cardSelection []string) *domain.Deck {
	var cards []*domain.Card

	if len(cardSelection) < 1 && !shuffle {
		cards = CreateCard()
	} else if len(cardSelection) < 1 && shuffle {
		newCards := CreateCard()
		cards = ShuffleCard(newCards)
	} else if len(cardSelection) > 0 && shuffle {
		newCard := CreateCard()
		selectedCards := MakeCardSelection(newCard, cardSelection)
		cards = ShuffleCard(selectedCards)
	} else {
		newCard := CreateCard()
		cards = MakeCardSelection(newCard, cardSelection)
	}

	deck = &domain.Deck{
		ID:        uuid.New().String(),
		Shuffled:  shuffle,
		Remaining: uint8(len(cards)),
		Cards:     cards,
	}

	Decks = append(Decks, deck)

	return deck
}

// CreateCard creates an array of 52 standard card comprising of the 
// following suits eg: diamonds, clubs, hearts and spades 
func CreateCard() []*domain.Card {
	var code string
	var cards []*domain.Card
	for _, suit := range suits {
		for _, value := range values {
			if len(value) > 2 {
				code = fmt.Sprintf("%s%s", string(value[0]), string(suit[0]))
			} else {
				code = fmt.Sprintf("%s%s", string(value), string(suit[0]))
			}
			card := &domain.Card{Value: value, Suit: suit, Code: code}
			cards = append(cards, card)
		}
	}

	return cards
}

// MakeCardSelection creates an array of standard card comprising of a 
// specified selection 
func MakeCardSelection(cards []*domain.Card, selection []string) []*domain.Card {
	var selectedCards []*domain.Card
	for _, card := range cards {
		for _, selection := range selection {
			if card.Code == selection {
				selectedCards = append(selectedCards, card)
			}
		}
	}

	return selectedCards
}

// ShuffleCard shuffles the array of cards 
func ShuffleCard(cards []*domain.Card) []*domain.Card {
	for i := len(cards) - 1; i > 0; i-- {
		rand.Seed(time.Now().UnixNano())
		newIndex := rand.Intn(i)
		prevCard := cards[newIndex]
		cards[newIndex] = cards[i]
		cards[i] = prevCard
	}

	return cards
}
