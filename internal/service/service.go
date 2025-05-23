package service

import (
	"github.com/TimNikolaev/drag-chat/internal/models"
	"github.com/TimNikolaev/drag-chat/internal/repository"
	"github.com/go-redis/redis/v8"
)

type Authorization interface {
}

type Chat interface {
}

type Chatting interface {
	GetChats(userID uint64) ([]models.Chat, error)
	GetHistory(string) ([]string, error)
	Publish(message models.Message) error
	Subscribe(chatIDs []string) *redis.PubSub
}

type Service struct {
	Authorization
	Chat
	Chatting
}

func New(repository *repository.Repository, rClient *redis.Client) *Service {
	return &Service{
		Chatting: NewChattingService(repository.Chatting, rClient),
	}
}
