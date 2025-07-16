package repository

import (
	"github.com/TimNikolaev/drag-chat/internal/models"
	"github.com/TimNikolaev/drag-chat/internal/repository/postgres"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user *models.User) (uint, error)
	GetUser(email, password string) (*models.User, error)
}

type Chat interface {
	CreateChat(userID uint, companionIDs []uint, chatName string) (*models.Chat, error)
	GetUserIDByUserName(userName string) (uint, error)
	GetChats(userID uint) ([]models.Chat, error)
	CreateMessage(chatID uint, senderID uint, text string) (*models.Message, error)
	GetMessages(userID uint, chatID uint) ([]models.Message, error)
	DeleteMessage(userID, chatID uint, messageID uint64) error
}

type Repository struct {
	Authorization
	Chat
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: postgres.NewAuthRepository(db),
		Chat:          postgres.NewChatRepository(db),
	}
}
