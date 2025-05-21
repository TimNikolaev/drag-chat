package repository

import (
	"github.com/TimNikolaev/drag-chat/internal/models"
	"github.com/TimNikolaev/drag-chat/internal/repository/postgres"
)

type Authorization interface {
}

type Chat interface {
	CreateChat() *models.Chat
}

type Chatting interface {
	GetChats(userID uint64) ([]models.Chat, error)
}

type Repository struct {
	Authorization
	Chat
	Chatting
}

func New() *Repository {
	return &Repository{
		Chatting: postgres.NewChattingRepository(),
	}
}
