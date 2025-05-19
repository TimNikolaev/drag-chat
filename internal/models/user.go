package models

type User struct {
	ID       uint   `json:"-"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
