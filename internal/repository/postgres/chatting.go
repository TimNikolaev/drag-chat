package postgres

import "github.com/TimNikolaev/drag-chat/internal/models"

type ChattingRepository struct {
}

func NewChattingRepository() *ChattingRepository {
	return &ChattingRepository{}
}

func (r *ChattingRepository) GetChats(userID uint64) ([]models.Chat, error) {
	return nil, nil
}
