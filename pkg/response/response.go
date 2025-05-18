package response

import (
	"log"

	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Message string `json:"message"`
}

func NewError(c *gin.Context, statusCode int, message string) {
	log.Fatalln(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{Message: message})
}
