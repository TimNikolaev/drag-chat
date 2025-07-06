package service

import (
	"github.com/TimNikolaev/drag-chat/internal/models"
	"github.com/TimNikolaev/drag-chat/internal/repository"
)

type ChatService struct {
	repository.Chat
}

func NewChatService(repo repository.Chat) *ChatService {
	return &ChatService{Chat: repo}
}

func (s *ChatService) CreateChat(userID uint, companionUserName string) (*models.Chat, error) {
	return nil, nil
}
