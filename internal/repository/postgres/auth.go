package postgres

import (
	"fmt"

	"github.com/TimNikolaev/drag-chat/internal/models"
	"github.com/jmoiron/sqlx"
)

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (r *AuthRepository) CreateUser(user *models.User) (uint, error) {
	var userID uint

	query := fmt.Sprintf("INSERT INTO %s (name, username, email ,password_hash) values ($1, $2, $3, $4) RETURNING id", usersTable)

	if err := r.db.DB.QueryRow(query, user.Name, user.UserName, user.Email, user.Password).Scan(&userID); err != nil {
		return 0, err
	}

	return userID, nil
}

func (r *AuthRepository) GetUser(email, password string) (*models.User, error) {
	return nil, nil
}
