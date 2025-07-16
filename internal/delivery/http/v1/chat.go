package v1

import (
	"net/http"
	"strconv"

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
	userID, err := GetUserID(c)
	if err != nil {
		return
	}

	chatID, err := strconv.Atoi(c.Param("chat_id"))
	if err != nil {
		response.NewError(c, http.StatusBadRequest, err.Error())
		return
	}

	messages, err := h.chatService.GetMessages(uint(userID), uint(chatID))
	if err != nil {
		response.NewError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, models.GetMessagesResponse{
		Data: messages,
	})
}

func (h *Handler) UpdateMessage(c *gin.Context) {

}

func (h *Handler) DeleteMessage(c *gin.Context) {
	userID, err := GetUserID(c)
	if err != nil {
		return
	}

	chatID, err := strconv.Atoi(c.Param("chat_id"))
	if err != nil {
		response.NewError(c, http.StatusBadRequest, err.Error())
		return
	}

	messageID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.NewError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.chatService.DeleteMessage(uint(userID), uint(chatID), uint64(messageID)); err != nil {
		response.NewError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response.StatusResponse{Status: "ok"})
}
