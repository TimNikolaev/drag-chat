package models

type Message struct {
	ID uint
	Chat
	Text string
	Time int64
}
