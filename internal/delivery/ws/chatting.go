package ws

import (
	"net/http"

	"github.com/TimNikolaev/drag-chat/pkg/response"
	"github.com/gin-gonic/gin"
)

func (ws *WSHandler) Connecting(c *gin.Context) {
	conn, err := ws.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		response.NewError(c, http.StatusBadRequest, err.Error())
		return
	}
	defer conn.Close()
}
