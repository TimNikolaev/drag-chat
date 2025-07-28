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
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

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

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &chat, nil
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
	var messagesCount uint64

	tx, err := r.db.Beginx()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	queryGetCount := fmt.Sprintf("SELECT c.messages_count FROM %s c JOIN %s uc ON c.id = uc.chat_id WHERE uc.user_id=$1 AND uc.chat_id=$2 FOR UPDATE", chatsTable, usersChatsTable)

	if err := tx.Get(&messagesCount, queryGetCount, senderID, chatID); err != nil {
		return nil, err
	}

	var message models.Message

	queryCreateMessage := fmt.Sprintf("INSERT INTO %s (id, chat_id, sender_id, text_body) VALUES ($1, $2, $3, $4) RETURNING *", messagesTable)

	if err := tx.QueryRow(queryCreateMessage, messagesCount+1, chatID, senderID, text).Scan(&message); err != nil {
		return nil, err
	}

	queryIncrCount := fmt.Sprintf("UPDATE SET %s messages_count = messages_count+1 WHERE id=$1", chatsTable)

	if _, err := tx.Exec(queryIncrCount, chatID); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
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
		return nil, fmt.Errorf("user %d is not a member of chat %d", userID, chatID)
	}

	var messages []models.Message

	queryGetMessages := fmt.Sprintf("SELECT * FROM %s WHERE chat_id=$1", messagesTable)

	if err := r.db.Select(&messages, queryGetMessages, chatID); err != nil {
		return nil, err
	}

	return messages, nil
}

func (r *ChatRepository) UpdateMessage(chatID uint, messageID uint64, userID uint, text string) (*models.Message, error) {
	var exists bool

	queryExists := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s c JOIN %s uc ON c.id = uc.chat_id WHERE uc.user_id=$1 AND uc.chat_id=$2)", chatsTable, usersChatsTable)

	if err := r.db.QueryRow(queryExists, userID, chatID).Scan(&exists); err != nil {
		return nil, err
	}

	if !exists {
		return nil, fmt.Errorf("user %d is not a member of chat %d", userID, chatID)
	}

	var updatedMessage models.Message

	queryUpdateMessage := fmt.Sprintf("UPDATE %s SET text_body = $1, is_edited = 'true' WHERE sender_id = $2 AND chat_id = $3 AND id = $4", messagesTable)

	if err := r.db.QueryRow(queryUpdateMessage, text, userID, chatID, messageID).Scan(&updatedMessage); err != nil {
		return nil, err
	}

	return &updatedMessage, nil
}

func (r *ChatRepository) DeleteMessage(userID, chatID uint, messageID uint64) error {
	var exists bool

	queryExists := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s c JOIN %s uc ON c.id = uc.chat_id WHERE uc.user_id=$1 AND uc.chat_id=$2)", chatsTable, usersChatsTable)

	if err := r.db.QueryRow(queryExists, userID, chatID).Scan(&exists); err != nil {
		return err
	}

	if !exists {
		return fmt.Errorf("user %d is not a member of chat %d", userID, chatID)
	}

	queryDeleteMessage := fmt.Sprintf("DELETE FROM %s WHERE id=$1 AND chat_id=$2 AND sender_id=$3", messagesTable)

	_, err := r.db.Exec(queryDeleteMessage, messageID, chatID, userID)

	return err
}
