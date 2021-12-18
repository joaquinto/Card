package handler

import "card/services"

type handler struct {
	cardService *services.CardService
}

func NewHandler(cardService *services.CardService) *handler {
	return &handler{cardService}
}