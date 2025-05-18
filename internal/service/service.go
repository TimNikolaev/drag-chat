package service

import "github.com/TimNikolaev/drag-chat/internal/repository"

type Service struct {
	repository.AuthRepository
	repository.ChatRepository
}

func New(repository *repository.Repository) *Service {
	return &Service{
		AuthRepository: repository.AuthRepository,
		ChatRepository: repository.ChatRepository,
	}
}
