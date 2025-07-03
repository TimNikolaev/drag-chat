package delivery

import (
	v1 "github.com/TimNikolaev/drag-chat/internal/delivery/v1"
	"github.com/TimNikolaev/drag-chat/internal/service"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	handler_v1 *v1.Handler
}

func New(service *service.Service) *Controller {
	return &Controller{
		handler_v1: v1.New(service),
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
				auth.POST("/sign-up", c.handler_v1.SignUp)
				auth.POST("/sign-in", c.handler_v1.SignIn)
			}

			chats := v1.Group("/chats", c.handler_v1.UserIdentity())
			{
				chats.POST("/", c.handler_v1.CreateChat)
				chats.GET("/", c.handler_v1.GetChats)

				messages := chats.Group(":id/messages")
				{
					messages.GET("/", c.handler_v1.GetMessages)
					messages.PUT("/:id", c.handler_v1.UpdateMessage)
					messages.DELETE("/:id", c.handler_v1.DeleteMessage)
				}
			}
		}

		//v....
	}

	return router
}

func (c *Controller) InitWSRouts() *gin.Engine {
	router := gin.New()
	router.GET("/ws", c.handler_v1.UserIdentity(), c.handler_v1.Chatting)

	return router
}
