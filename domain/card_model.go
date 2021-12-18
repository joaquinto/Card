package domain

// Card This is the card struct
type Card struct {
	Value string `json:"value"`
	Suit  string `json:"suit"`
	Code  string `json:"code"`
}

// Deck This is the deck struct
type Deck struct {
	ID        string  `json:"deck_id,omitempty"`
	Shuffled  bool    `json:"shuffled"`
	Remaining uint8   `json:"remaining"`
	Cards     []*Card `json:"cards,omitempty"`
}

