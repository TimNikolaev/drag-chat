package models

type Chat struct {
	ID       uint   `json:"chat_id" db:"id"`
	ChatName string `json:"chat_name" db:"chat_name"`
}

type CreateChatRequest struct {
	CompanionUserNames []string `json:"companion_user_names"`
	ChatName           string   `json:"chat_name"`
}
