package handler

import (
	"card/domain"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

var cardCode map[string]string = map[string]string{
	"AC": "AC", "2C": "2C", "3C": "3C", "4C": "4C", "5C": "5C",
	"6C": "6C", "7C": "7C", "8C": "8C", "9C": "9C", "10C": "10C",
	"JC": "JC", "QC": "QC", "KC": "KC", "AD": "AD", "2D": "2D",
	"3D": "3D", "4D": "4D", "5D": "5D", "6D": "6D", "7D": "7D",
	"8D": "8D", "9D": "9D", "10D": "10D", "JD": "JD", "QD": "QD",
	"KD": "KD", "AH": "AH", "2H": "2H", "3H": "3H", "4H": "4H",
	"5H": "5H", "6H": "6H", "7H": "7H", "8H": "8H", "9H": "9H",
	"10H": "10H", "JH": "JH", "QH": "QH", "KH": "KH", "AS": "AS",
	"2S": "2S", "3S": "3S", "4S": "4S", "5S": "5S", "6S": "6S",
	"7S": "7S", "8S": "8S", "9S": "9S", "10S": "10S", "JS": "JS",
	"QS": "QS", "KS": "KS",
}

// CreateDeck This implements the create deck handler
// This implementation also handles edge cases for when
// the shuffled and cards query parameters are set or not
// by default without passing any query parameter,
// this generates the standard un-shuffled 52 deck of card
func (h *handler) CreateDeck(ctx echo.Context) error {
	var shuffled bool
	var cards []string
	shuffle := ctx.QueryParam("shuffled")
	selectedCards := ctx.QueryParam("cards")

	if shuffle == "true" && shuffle != "" {
		shuffled = true
	} else if shuffle == "false" || shuffle == "" {
		shuffled = false
	} else {
		return ctx.JSON(http.StatusBadRequest, "Shuffle should only be true or false")
	}

	if selectedCards != "" {
		cards = strings.Split(selectedCards, ",")

		for _, code := range cards {
			if cardCode[code] == "" {
				return ctx.JSON(http.StatusBadRequest, "Invalid card code, please input a valid card code")
			}
		}
	}

	deck := h.cardService.CreateDeck(shuffled, cards)

	return ctx.JSON(http.StatusCreated, domain.Deck{
			ID:        deck.ID,
			Shuffled:  deck.Shuffled,
			Remaining: deck.Remaining,
		},
	)
}
