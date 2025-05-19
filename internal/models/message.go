package models

import "time"

type Message struct {
	ID uint `json:"message_id"`
	Chat
	User
	Text string    `json:"text"`
	Time time.Time `json:"time"`
}
