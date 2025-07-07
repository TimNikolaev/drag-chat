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
	CreateChat(userID uint, companionID []uint, chatName string) (*models.Chat, error)
	GetUserByUserName(userName string) (*models.User, error)
}

type Chatting interface {
	GetChats(userID uint) ([]models.Chat, error)
}

type Repository struct {
	Authorization
	Chat
	Chatting
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: postgres.NewAuthRepository(db),
		Chat:          postgres.NewChatRepository(db),
		Chatting:      postgres.NewChattingRepository(db),
	}
}
