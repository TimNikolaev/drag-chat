package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/TimNikolaev/drag-chat/internal/config"
	"github.com/TimNikolaev/drag-chat/internal/models"
	"github.com/TimNikolaev/drag-chat/internal/repository"
	"github.com/TimNikolaev/drag-chat/pkg/jwt"
)

type AuthService struct {
	auth   repository.Authorization
	config *config.Auth
}

func NewAuthService(repo repository.Authorization, cfg *config.Auth) *AuthService {
	return &AuthService{auth: repo, config: cfg}
}

const salt = "qwerty123456789"

func (s *AuthService) CreateUser(user *models.User) (uint, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.auth.CreateUser(user)
}

func (s *AuthService) GenerateToken(email, password string) (string, error) {
	user, err := s.auth.GetUser(email, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	return jwt.NewJWT(s.config.Secret).Generate(user.ID)
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	return jwt.NewJWT(s.config.Secret).Parse(accessToken)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
