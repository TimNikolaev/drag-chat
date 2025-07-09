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

	if len(companionUserNames) != 0 {
		for _, uname := range companionUserNames {
			id, err := s.chatRepository.GetUserIDByUserName(uname)
			if err != nil {
				return nil, err
			}
			companionIDs = append(companionIDs, id)
		}
	}

	return s.chatRepository.CreateChat(userID, companionIDs, chatName)
}
