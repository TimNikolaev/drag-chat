package models

type User struct {
	ID       uint   `json:"-" db:"id"`
	Name     string `json:"name"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
