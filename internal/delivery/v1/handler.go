package v1

import (
	"net/http"

	"github.com/TimNikolaev/drag-chat/internal/service"
	"github.com/gorilla/websocket"
)

type Handler struct {
	authService     service.Authorization
	chatService     service.Chat
	chattingService service.Chatting
	upgrader        *websocket.Upgrader
}

func New(service *service.Service) *Handler {
	return &Handler{
		authService:     service.Authorization,
		chatService:     service.Chat,
		chattingService: service.Chatting,
		upgrader: &websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin:     func(r *http.Request) bool { return true },
		},
	}
}
