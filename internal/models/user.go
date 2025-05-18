package models

type User struct {
	ID       uint
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
