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

func (s *ChatService) GetChats(userID uint) ([]models.Chat, error) {
	return s.chatRepository.GetChats(userID)
}

func (s *ChatService) GetMessages(userID uint, chatID uint) ([]models.Message, error) {
	return s.chatRepository.GetMessages(userID, chatID)
}

func (s *ChatService) UpdateMessage(chatID uint, messageID uint64, userID uint, text string) (*models.Message, error) {
	return nil, nil
}

func (s *ChatService) DeleteMessage(userID, chatID uint, messageID uint64) error {
	return s.chatRepository.DeleteMessage(userID, chatID, messageID)
}
