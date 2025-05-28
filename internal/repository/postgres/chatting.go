package postgres

import (
	"github.com/TimNikolaev/drag-chat/internal/models"
	"github.com/jmoiron/sqlx"
)

type ChattingRepository struct {
	db *sqlx.DB
}

func NewChattingRepository(db *sqlx.DB) *ChattingRepository {
	return &ChattingRepository{db: db}
}

func (r *ChattingRepository) GetChats(userID uint) ([]models.Chat, error) {
	return nil, nil
}
