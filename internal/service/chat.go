package service

import (
	"github.com/TimNikolaev/drag-chat/internal/models"
	"github.com/TimNikolaev/drag-chat/internal/repository"
)

type ChatService struct {
	chatRepository repository.Chat
}

func NewChatService(repo repository.Chat) *ChatService {
	return &ChatService{chatRepository: repo}
}

func (s *ChatService) CreateChat(userID uint, companionUserNames []string, chatName string) (*models.Chat, error) {
	var companionIDs []uint
	for _, name := range companionUserNames {
		companion, err := s.chatRepository.GetUserByUserName(name)
		if err != nil {
			return nil, err
		}
		companionIDs = append(companionIDs, companion.ID)
	}
	return s.chatRepository.CreateChat(userID, companionIDs, chatName)
}
