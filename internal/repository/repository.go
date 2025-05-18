package repository

import "github.com/TimNikolaev/drag-chat/internal/models"

type AuthRepository interface {
}

type ChatRepository interface {
	CreateChat() *models.Chat
}

type Repository struct {
	AuthRepository
	ChatRepository
}

func New() *Repository {
	return &Repository{}
}
