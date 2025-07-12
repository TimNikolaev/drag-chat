package service

import (
	"context"
	"encoding/json"

	"github.com/TimNikolaev/drag-chat/internal/models"
	"github.com/TimNikolaev/drag-chat/internal/repository"
	"github.com/go-redis/redis/v8"
)

type ChattingService struct {
	chatRepository repository.Chat
	rClient        *redis.Client
}

func NewChattingService(repo repository.Chat, rClient *redis.Client) *ChattingService {
	return &ChattingService{rClient: rClient}
}

var ctx = context.Background()

func (s *ChattingService) GetChats(userID uint) ([]models.Chat, error) {
	return s.chatRepository.GetChats(userID)
}

func (s *ChattingService) GetHistory(chatID string) ([]string, error) {
	return s.rClient.LRange(ctx, chatID, 0, -1).Result()
}

func (s *ChattingService) Publish(msg *models.Message) error {
	msgJSON, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	if err := s.rClient.Publish(ctx, string(rune(msg.ChatID)), msgJSON).Err(); err != nil {
		return err
	}

	s.rClient.LPush(ctx, string(rune(msg.ChatID)), msgJSON)

	return nil
}

func (s *ChattingService) Subscribe(chatIDs []string) *redis.PubSub {
	return s.rClient.Subscribe(ctx, chatIDs...)
}
