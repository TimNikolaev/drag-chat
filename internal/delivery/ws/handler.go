package ws

import (
	"github.com/TimNikolaev/drag-chat/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WSHandler struct {
	service.Chatting
	upgrader *websocket.Upgrader
}

func NewWSHandler(service *service.Service, upgrader *websocket.Upgrader) *WSHandler {
	return &WSHandler{
		Chatting: service.Chatting,
		upgrader: upgrader,
	}
}

func (ws *WSHandler) InitConnectRout() *gin.Engine {
	router := gin.New()
	router.GET("/ws", ws.Connecting)

	return router
}
