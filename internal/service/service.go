package service

import (
	"github.com/TimNikolaev/drag-chat/internal/models"
	"github.com/TimNikolaev/drag-chat/internal/repository"
)

type Authorization interface {
}

type Chat interface {
}

type Chatting interface {
	Publish(chatID int64, message models.Message) error
	Subscribe(chatIDs ...int64) error
}

type Service struct {
	Authorization
	Chat
	Chatting
}

func New(repository *repository.Repository) *Service {
	return &Service{
		Chatting: NewChattingService(repository.Chatting),
	}
}
