package service

import (
	"github.com/TimNikolaev/drag-chat/internal/config"
	"github.com/TimNikolaev/drag-chat/internal/models"
	"github.com/TimNikolaev/drag-chat/internal/repository"
	"github.com/go-redis/redis/v8"
)

type Authorization interface {
	CreateUser(user *models.User) (uint, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type Chat interface {
	CreateChat(userID uint, companionUserNames []string, chatName string) (*models.Chat, error)
	GetChats(userID uint) ([]models.Chat, error)
	GetMessages(userID uint, chatID uint) ([]models.Message, error)
	UpdateMessage(chatID uint, messageID uint64, userID uint, text string) (*models.Message, error)
	DeleteMessage(userID, chatID uint, messageID uint64) error
}

type Chatting interface {
	GetChats(userID uint) ([]models.Chat, error)
	GetHistory(string) ([]string, error)
	CreateMessage(chatID, senderID uint, text string) (*models.Message, error)
	Publish(message *models.Message) error
	Subscribe(chatIDs []string) *redis.PubSub
}

type Service struct {
	Authorization
	Chat
	Chatting
}

func New(repository *repository.Repository, rClient *redis.Client, cfg *config.Auth) *Service {
	return &Service{
		Authorization: NewAuthService(repository.Authorization, cfg),
		Chat:          NewChatService(repository.Chat),
		Chatting:      NewChattingService(repository.Chat, rClient),
	}
}
