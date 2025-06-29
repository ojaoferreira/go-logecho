package service

import (
	"go-logecho/model"
	"go-logecho/repository"
)

type MessageService interface {
	CreateMessage(message model.Message)
}

type messageService struct {
	repo repository.MessageRepository
}

func NewMessageService(r repository.MessageRepository) MessageService {
	return &messageService{r}
}

func (s *messageService) CreateMessage(message model.Message) {
	s.repo.CreateMessage(message)
}
