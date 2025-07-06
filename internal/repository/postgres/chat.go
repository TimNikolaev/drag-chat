package postgres

import (
	"github.com/TimNikolaev/drag-chat/internal/models"
	"github.com/jmoiron/sqlx"
)

type ChatRepository struct {
	db *sqlx.DB
}

func NewChatRepository(db *sqlx.DB) *ChatRepository {
	return &ChatRepository{
		db: db,
	}
}

func (r *ChatRepository) CreateChat(userID, companionID uint) (*models.Chat, error) {
	return nil, nil
}

func (r *ChatRepository) GetChatIDByUserName(userName string) (uint, error) {
	return 0, nil
}
