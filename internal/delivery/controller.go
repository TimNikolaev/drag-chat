package delivery

import (
	"github.com/TimNikolaev/drag-chat/internal/delivery/handlers"
	"github.com/TimNikolaev/drag-chat/internal/service"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	handler *handlers.Handler
}

func New(service *service.Service) *Controller {
	return &Controller{
		handler: handlers.New(service),
	}
}

func (c *Controller) InitRouts() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/sign-up", c.handler.SignUp)
			auth.POST("/sign-in", c.handler.SignIn)
		}
		chats := api.Group("/chats")
		{
			chats.POST("/", c.handler.CreateChat)
			chats.GET("/", c.handler.GetChats)

			messages := chats.Group(":id/messages")
			{
				messages.GET("/", c.handler.GetMessages)
				messages.PUT("/:id", c.handler.UpdateMessage)
				messages.DELETE("/:id", c.handler.DeleteMessage)
			}
		}
	}

	return router
}

func (c *Controller) InitWSRouts() *gin.Engine {
	router := gin.New()
	router.GET("/ws", c.handler.Chatting)

	return router
}
