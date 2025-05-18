package service

import "github.com/TimNikolaev/drag-chat/internal/repository"

type AuthService struct {
	repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) *AuthService {
	return &AuthService{AuthRepository: repo}
}
