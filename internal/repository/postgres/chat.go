package postgres

import (
	"errors"
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

	queryCreateChat := fmt.Sprintf("INSERT INTO %s (chat_name) values ($1) RETURNING *", chatsTable)

	if err := tx.QueryRow(queryCreateChat, chatName).Scan(&chat); err != nil {
		return nil, err
	}

	queryCreateUsersChats := fmt.Sprintf("INSERT INTO %s (user_id, chat_id) values ($1, $2)", usersChatsTable)

	for _, id := range companionIDs {
		if _, err := tx.Exec(queryCreateUsersChats, id, chat.ID); err != nil {
			return nil, err
		}
	}

	if _, err := tx.Exec(queryCreateUsersChats, userID, chat.ID); err != nil {
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
	var chats []models.Chat

	query := fmt.Sprintf("SELECT c.id, c.name FROM %s c JOIN %s uc ON c.id = uc.chat_id WHERE uc.user_id=$1", chatsTable, usersChatsTable)

	if err := r.db.Select(&chats, query, userID); err != nil {
		return nil, err
	}

	return chats, nil
}

func (r *ChatRepository) CreateMessage(chatID uint, senderID uint, text string) (*models.Message, error) {
	var message models.Message

	query := fmt.Sprintf("INSERT INTO %s (chat_id, sender_id, text_body) VALUES ($1, $2, $3) RETURNING *", messagesTable)

	if err := r.db.QueryRow(query, chatID, senderID, text).Scan(&message); err != nil {
		return nil, err
	}

	return &message, nil
}

func (r *ChatRepository) GetMessages(userID uint, chatID uint) ([]models.Message, error) {
	var exists bool

	queryExists := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s c JOIN %s uc ON c.id = uc.chat_id WHERE uc.user_id=$1 AND uc.chat_id=$2)", chatsTable, usersChatsTable)

	if err := r.db.QueryRow(queryExists, userID, chatID).Scan(&exists); err != nil {
		return nil, err
	}

	if !exists {
		return nil, errors.New("row not exists")
	}

	var messages []models.Message

	queryGetMessages := fmt.Sprintf("SELECT * FROM %s WHERE chat_id=$1", messagesTable)

	if err := r.db.Select(&messages, queryGetMessages, chatID); err != nil {
		return nil, err
	}

	return messages, nil
}
