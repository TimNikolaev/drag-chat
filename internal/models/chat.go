package models

type Chat struct {
	ID       uint   `json:"chat_id"`
	ChatName string `json:"chat_name"`
}

type RequestPersonalChat struct {
	CompanionUserName string `json:"companion_user_name"`
}
