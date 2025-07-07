package models

type Chat struct {
	ID       uint   `json:"chat_id"`
	ChatName string `json:"chat_name"`
}

type CreateChatRequest struct {
	CompanionUserNames []string `json:"companion_user_names"`
	ChatName           string   `json:"chat_name"`
}
