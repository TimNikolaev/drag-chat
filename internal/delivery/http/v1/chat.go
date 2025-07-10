package v1

import (
	"net/http"

	"github.com/TimNikolaev/drag-chat/internal/models"
	"github.com/TimNikolaev/drag-chat/pkg/response"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateChat(c *gin.Context) {
	userID, err := GetUserID(c)
	if err != nil {
		return
	}

	var input models.CreateChatRequest

	if err := c.BindJSON(&input); err != nil {
		response.NewError(c, http.StatusInternalServerError, err.Error())
		return
	}

	chat, err := h.chatService.CreateChat(uint(userID), input.CompanionUserNames, input.ChatName)
	if err != nil {
		response.NewError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]any{"chat": chat})

}

func (h *Handler) GetChats(c *gin.Context) {
	userID, err := GetUserID(c)
	if err != nil {
		return
	}

	chats, err := h.chatService.GetChats(uint(userID))
	if err != nil {
		response.NewError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, models.GetChatsResponse{
		Data: chats,
	})
}

func (h *Handler) GetMessages(c *gin.Context) {

}

func (h *Handler) UpdateMessage(c *gin.Context) {

}

func (h *Handler) DeleteMessage(c *gin.Context) {

}
