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

	var input models.RequestPersonalChat

	if err := c.BindJSON(&input); err != nil {
		response.NewError(c, http.StatusInternalServerError, err.Error())
		return
	}

	chat, err := h.chatService.CreateChat(uint(userID), input.CompanionUserName)
	if err != nil {
		response.NewError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(200, chat)

}

func (h *Handler) GetChats(c *gin.Context) {

}

func (h *Handler) GetMessages(c *gin.Context) {

}

func (h *Handler) UpdateMessage(c *gin.Context) {

}

func (h *Handler) DeleteMessage(c *gin.Context) {

}
