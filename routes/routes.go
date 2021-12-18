package routes

import (
	"github.com/labstack/echo/v4"

	"card/handler"
	"card/services"
)

func Run(e *echo.Echo) *echo.Echo {
	cs := services.NewCardService()
	h:= handler.NewHandler(cs)

	e.POST("/api/v1/decks", h.CreateDeck)

	e.GET("/api/v1/decks/:deckId", h.OpenDeck)

	e.PATCH("/api/v1/decks/:deckId/draw-card", h.DrawCard)

	e.Any("*", h.RouteNotFound)

	return e
}
