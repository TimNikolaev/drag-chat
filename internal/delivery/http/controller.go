package http

import (
	v1 "github.com/TimNikolaev/drag-chat/internal/delivery/http/v1"
	"github.com/TimNikolaev/drag-chat/internal/service"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	handlerV1 *v1.Handler
}

func New(service *service.Service) *Controller {
	return &Controller{
		handlerV1: v1.New(service),
	}
}

func (c *Controller) InitRouts() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		//v1
		v1 := api.Group("/v1")
		{
			auth := v1.Group("/auth")
			{
				auth.POST("/sign-up", c.handlerV1.SignUp)
				auth.POST("/sign-in", c.handlerV1.SignIn)
			}

			chats := v1.Group("/chats", c.handlerV1.UserIdentity())
			{
				chats.GET("/", c.handlerV1.GetChats)
				chats.POST("/", c.handlerV1.CreateChat)

				messages := chats.Group(":chat_id/messages")
				{
					messages.GET("/", c.handlerV1.GetMessages)
					messages.PUT("/:id", c.handlerV1.UpdateMessage) //
					messages.DELETE("/:id", c.handlerV1.DeleteMessage)
				}
			}
			v1.GET("/ws/chatting", c.handlerV1.UserIdentity(), c.handlerV1.Chatting)
		}

		//v....
	}

	return router
}
