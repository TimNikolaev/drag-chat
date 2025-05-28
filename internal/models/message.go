package models

import "time"

type Message struct {
	ID       uint      `json:"message_id"`
	ChatID   uint      `json:"chat_id"`
	UserID   uint      `json:"user_id"`
	Text     string    `json:"text"`
	SendTime time.Time `json:"send_time"`
}
