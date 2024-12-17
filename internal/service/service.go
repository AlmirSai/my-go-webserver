package service

import (
	"my-go-webserver/internal/storage"
)

type Service interface {
	GetWelcomeMessage() string
}

type service struct {
	storage storage.Storage
}

func NewService(s storage.Storage) Service {
	return &service{storage: s}
}

func (s *service) GetWelcomeMessage() string {
	return s.storage.GetMessage()
}
