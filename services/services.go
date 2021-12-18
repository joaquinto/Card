package services

import "card/domain"

type CardService struct{}

func NewCardService() *CardService {
	return &CardService{}
}

var _ domain.CardService = &CardService{}
