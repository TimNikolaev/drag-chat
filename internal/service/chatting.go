package service

import (
	"github.com/TimNikolaev/drag-chat/internal/models"
	"github.com/TimNikolaev/drag-chat/internal/repository"
)

type ChattingService struct {
	repository.Chatting
}

func NewChattingService(repo repository.Chatting) *ChattingService {
	return &ChattingService{Chatting: repo}
}

func (ch *ChattingService) Publish(chatID int64, message models.Message) error {
	return nil
}

func (ch *ChattingService) Subscribe(chatIDs ...int64) error {
	return nil
}
