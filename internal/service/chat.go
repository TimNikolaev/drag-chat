package service

import "github.com/TimNikolaev/drag-chat/internal/repository"

type ChatService struct {
	repository.Chat
}

func NewChatService(repo repository.Chat) *ChatService {
	return &ChatService{Chat: repo}
}
