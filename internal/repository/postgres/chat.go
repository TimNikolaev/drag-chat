package postgres

import (
	"fmt"

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

func (r *ChatRepository) CreateChat(userID uint, companionIDs []uint, chatName string) (*models.Chat, error) {
	var chat models.Chat

	tx, err := r.db.Beginx()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	query := fmt.Sprintf("INSERT INTO %s (chat_name) values ($1) RETURNING *", chatsTable)
	if err := tx.QueryRow(query, chatName).Scan(&chat); err != nil {
		return nil, err
	}

	if companionIDs != nil {
		query = fmt.Sprintf("INSERT INTO %s (user_id, chat_id) values ($1, $2)", usersChatsTable)
		for _, id := range companionIDs {
			if _, err := tx.Exec(query, id, chat.ID); err != nil {
				return nil, err
			}
		}
	}

	if _, err := tx.Exec(query, userID, chat.ID); err != nil {
		return nil, err
	}

	return &chat, tx.Commit()
}

func (r *ChatRepository) GetUserIDByUserName(userName string) (uint, error) {
	var userID uint

	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1", usersTable)

	err := r.db.Get(&userID, query, userName)

	return userID, err
}

func (r *ChatRepository) GetChats(userID uint) ([]models.Chat, error) {
	return nil, nil
}
