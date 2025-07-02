package postgres

import (
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
	return 0, nil
}

func (r *AuthRepository) GetUser(email, password string) (*models.User, error) {
	return nil, nil
}
