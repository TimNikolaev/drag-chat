package v1

import (
	"net/http"

	"github.com/TimNikolaev/drag-chat/internal/models"
	"github.com/TimNikolaev/drag-chat/pkg/response"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SignUp(c *gin.Context) {
	var input models.User

	if err := c.BindJSON(&input); err != nil {
		response.NewError(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.authService.CreateUser(&input)
	if err != nil {
		response.NewError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]any{"id": id})

}

func (h *Handler) SignIn(c *gin.Context) {
	var input models.SignInRequest

	if err := c.BindJSON(&input); err != nil {
		response.NewError(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.authService.GenerateToken(input.Email, input.Password)
	if err != nil {
		response.NewError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]any{"token": token})
}
