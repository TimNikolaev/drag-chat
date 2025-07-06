package v1

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/TimNikolaev/drag-chat/internal/models"
	"github.com/TimNikolaev/drag-chat/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func (h *Handler) Chatting(c *gin.Context) {
	userID, err := GetUserID(c)
	if err != nil {
		return
	}

	conn, err := h.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		response.NewError(c, http.StatusBadRequest, err.Error())
		return
	}
	defer conn.Close()

	chats, err := h.chattingService.GetChats(uint(userID))
	if err != nil {
		//error handling and client notification using ws
		return
	}

	chatIDsString := []string{}
	for _, chat := range chats {
		chatIDsString = append(chatIDsString, string(rune(chat.ID)))
	}

	go h.sendMessages(conn)

	h.getHistory(conn, chatIDsString)

	h.getMessages(conn, chatIDsString)
}

func (h *Handler) sendMessages(conn *websocket.Conn) {
	for {
		var msgInput models.MessageRequest

		if err := conn.ReadJSON(&msgInput); err != nil {
			//error logging
			continue
		}

		msg := models.Message{
			ID:       msgInput.ID,
			ChatID:   msgInput.ChatID,
			UserID:   msgInput.UserID,
			Text:     msgInput.Text,
			SendTime: time.Now(),
		}

		if err := h.chattingService.Publish(&msg); err != nil {
			//error handling and client notification using ws
			continue
		}

	}

}

func (h *Handler) getHistory(conn *websocket.Conn, chatsIDs []string) {
	for _, chatID := range chatsIDs {
		historyMsgs, err := h.chattingService.GetHistory(chatID)
		if err == nil {
			for _, h := range historyMsgs {
				var msg models.Message
				if err := json.Unmarshal([]byte(h), &msg); err != nil {
					//error handling and client notification using ws
					continue
				}
				if err := conn.WriteJSON(msg); err != nil {
					// error logging
				}
			}

		}

	}

}

func (h *Handler) getMessages(conn *websocket.Conn, chatIDsString []string) {
	var message models.Message

	pubsub := h.chattingService.Subscribe(chatIDsString)

	defer pubsub.Close()

	chanMsgs := pubsub.Channel()

	for {
		msg := <-chanMsgs

		if err := json.Unmarshal([]byte(msg.Payload), &message); err != nil {
			//error handling and client notification using ws
			continue
		}

		if err := conn.WriteJSON(message); err != nil {
			// error logging
			break
		}

	}

}
