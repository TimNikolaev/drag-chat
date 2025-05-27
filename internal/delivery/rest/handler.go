package rest

import (
	"github.com/TimNikolaev/drag-chat/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service.Authorization
	service.Chat
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		Authorization: service.Authorization,
		Chat:          service.Chat,
	}
}

func (h *Handler) InitRouts() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/sign-up", h.signUp)
			auth.POST("/sign-in", h.signIn)
		}
		chats := api.Group("/chats")
		{
			chats.POST("/", h.createChat)
			chats.GET("/", h.getChats)

			messages := chats.Group(":id/messages")
			{
				messages.GET("/", h.getMessages)
			}
		}
	}

	return router
}
