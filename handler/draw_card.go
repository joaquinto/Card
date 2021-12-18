package handler

import (
	"card/domain"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// DrawCard This implements the draw card handler
func (h *handler) DrawCard(ctx echo.Context) error {
	deckId := ctx.Param("deckId")
	noOfCard := ctx.QueryParam("count")

	if deckId == "" {
		return ctx.JSON(http.StatusBadRequest, "Deck id is empty")
	}

	if noOfCard == "" {
		return ctx.JSON(http.StatusBadRequest, "Count is empty")
	}

	count, _ := strconv.Atoi(noOfCard)

	if count < 1 {
		return ctx.JSON(http.StatusBadRequest, "Count must me greater than or equal to 1")
	}

	card, err := h.cardService.DrawCard(deckId, uint8(count))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, err.Error())
	}

	return ctx.JSON(http.StatusOK, map[string][]*domain.Card{
		"cards": card,
	})
}
