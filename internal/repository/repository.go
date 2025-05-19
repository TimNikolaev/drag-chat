package repository

import "github.com/TimNikolaev/drag-chat/internal/models"

type Authorization interface {
}

type Chat interface {
	CreateChat() *models.Chat
}

type Chatting interface {
}

type Repository struct {
	Authorization
	Chat
	Chatting
}

func New() *Repository {
	return &Repository{}
}
