package models

type User struct {
	ID       uint   `json:"-" db:"id"`
	Name     string `json:"name" db:"name"`
	UserName string `json:"user_name" db:"username"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password_hash"`
}

type SignInRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
