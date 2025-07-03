package v1

import (
	"errors"
	"net/http"
	"strings"

	"github.com/TimNikolaev/drag-chat/pkg/response"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) UserIdentity() gin.HandlerFunc {
	return func(c *gin.Context) {
		authedHeader := c.GetHeader(authorizationHeader)
		if !strings.HasPrefix(authedHeader, "Bearer ") {
			response.NewError(c, http.StatusUnauthorized, "invalid auth header")
			return
		}

		token := strings.TrimPrefix(authedHeader, "Bearer ")
		userId, err := h.authService.ParseToken(token)
		if err != nil {
			response.NewError(c, http.StatusUnauthorized, err.Error())
			return
		}

		c.Set(userCtx, userId)
	}
}

func GetUserID(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		response.NewError(c, http.StatusInternalServerError, "user id not found")
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		response.NewError(c, http.StatusInternalServerError, "user id is of invalid type")
		return 0, errors.New("user id is of invalid type")
	}

	return idInt, nil
}
