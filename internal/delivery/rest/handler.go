package rest

import (
	"github.com/TimNikolaev/drag-chat/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler() *Handler {
	return &Handler{}
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
		}
	}

	return router
}
