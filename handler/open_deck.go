package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// OpenDeck This implements the open deck handler
func (h *handler) OpenDeck(ctx echo.Context) error {
	deckId := ctx.Param("deckId")
	if deckId == "" {
		return ctx.JSON(http.StatusBadRequest, "Deck id is empty")
	}

	deck, err := h.cardService.OpenDeck(deckId)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, err.Error())
	}

	return ctx.JSON(http.StatusOK, deck)
}
