package service

import "github.com/TimNikolaev/drag-chat/internal/repository"

type ChatService struct {
	repository.ChatRepository
}

func NewChatService(repo repository.ChatRepository) *ChatService {
	return &ChatService{ChatRepository: repo}
}
