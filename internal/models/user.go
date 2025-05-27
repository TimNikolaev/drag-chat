package models

type User struct {
	ID       uint64 `json:"-"`
	Name     string `json:"name"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
