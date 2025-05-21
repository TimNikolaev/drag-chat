package models

import "time"

type Message struct {
	ID     uint64    `json:"message_id"`
	ChatID uint64    `json:"chat_id"`
	UserID uint64    `json:"user_id"`
	Text   string    `json:"text"`
	Time   time.Time `json:"time"`
}
