package ws

import (
	"net/http"

	"github.com/TimNikolaev/drag-chat/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WSHandler struct {
	Chatting service.Chatting
	upgrader *websocket.Upgrader
}

func New(service *service.Service) *WSHandler {
	return &WSHandler{
		Chatting: service.Chatting,
		upgrader: &websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin:     func(r *http.Request) bool { return true },
		},
	}
}

func (ws *WSHandler) InitConnectRout() *gin.Engine {
	router := gin.New()
	router.GET("/ws", ws.chatting)

	return router
}
