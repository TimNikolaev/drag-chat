package models

import (
	"time"
)

type Message struct {
	ID       uint64    `json:"message_id" db:"id"`
	ChatID   uint      `json:"chat_id" db:"chat_id"`
	SenderID uint      `json:"user_id" db:"sender_id"`
	Text     string    `json:"text" db:"text_body"`
	IsEdited bool      `json:"is_edited" db:"is_edited"`
	SendTime time.Time `json:"send_time" db:"send_time"`
}

type SendMessageRequest struct {
	ChatID   uint   `json:"chat_id"`
	SenderID uint   `json:"user_id"`
	Text     string `json:"text"`
}

type GetMessagesResponse struct {
	Data []Message `json:"data"`
}
