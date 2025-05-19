package service

import "github.com/TimNikolaev/drag-chat/internal/repository"

type AuthService struct {
	repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{Authorization: repo}
}
