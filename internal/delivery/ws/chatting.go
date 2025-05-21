package ws

import (
	"net/http"
	"time"

	"github.com/TimNikolaev/drag-chat/internal/models"
	"github.com/TimNikolaev/drag-chat/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func (ws *WSHandler) chatting(c *gin.Context) {
	userID := 12
	conn, err := ws.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		response.NewError(c, http.StatusBadRequest, err.Error())
		return
	}
	defer conn.Close()

	chats, err := ws.Chatting.GetChats(uint64(userID))
	if err != nil {
		response.NewError(c, http.StatusInternalServerError, err.Error())
		return
	}

	chatsIDs := []uint64{}
	for _, chat := range chats {
		chatsIDs = append(chatsIDs, chat.ID)
	}

	ws.Chatting.Subscribe(chatsIDs...)

	go getMessages(ws, conn)
}

type messageRequest struct {
	ID     uint64 `json:"message_id"`
	ChatID uint64 `json:"chat_id"`
	UserID uint64 `json:"user_id"`
	Text   string `json:"text"`
}

func getMessages(ws *WSHandler, conn *websocket.Conn) error {
	for {
		var msgInput messageRequest

		if err := conn.ReadJSON(&msgInput); err != nil {
			return err
		}

		msg := models.Message{
			ID:     msgInput.ID,
			ChatID: msgInput.ChatID,
			UserID: msgInput.UserID,
			Text:   msgInput.Text,
			Time:   time.Now(),
		}

		if err := ws.Chatting.Publish(msg); err != nil {
			return err
		}

	}

}
