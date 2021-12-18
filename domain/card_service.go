package domain

// CardService This is an interface composition
// that needs to be implemented
type CardService interface {
	CreateDeck
	OpenDeck
	DrawCard
}

// CreateDeck This is the create deck interface
type CreateDeck interface {
	CreateDeck(shuffle bool, cardSelection []string) *Deck
}

// OpenDeck This is the open deck interface
type OpenDeck interface {
	OpenDeck(deckID string) (*Deck, error)
}

// DrawCard This is the draw card interface
type DrawCard interface {
	DrawCard(deckID string, noOfCard uint8) ([]*Card, error)
}