package service

import (
	"context"
	"encoding/json"

	"github.com/TimNikolaev/drag-chat/internal/models"
	"github.com/TimNikolaev/drag-chat/internal/repository"
	"github.com/go-redis/redis/v8"
)

type ChattingService struct {
	repository.Chatting
	rClient *redis.Client
}

func NewChattingService(repo repository.Chatting, rClient *redis.Client) *ChattingService {
	return &ChattingService{Chatting: repo}
}

var ctx = context.Background()

func (s *ChattingService) GetChats(userID uint) ([]models.Chat, error) {
	return s.Chatting.GetChats(userID)
}

func (s *ChattingService) GetHistory(chatID string) ([]string, error) {
	return s.rClient.LRange(ctx, chatID, 0, -1).Result()
}

func (s *ChattingService) Publish(msg *models.Message) error {
	msgByteJSON, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	if err := s.rClient.Publish(ctx, string(rune(msg.ChatID)), msgByteJSON).Err(); err != nil {
		return err
	}

	s.rClient.LPush(ctx, string(rune(msg.ChatID)), msgByteJSON)

	return nil
}

func (s *ChattingService) Subscribe(chatIDs []string) *redis.PubSub {
	return s.rClient.Subscribe(ctx, chatIDs...)
}
