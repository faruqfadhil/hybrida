package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func Success(c *gin.Context, data interface{}, msg string) {
	resp := Response{
		Data:    data,
		Message: msg,
	}

	c.JSON(http.StatusOK, resp)
}

func Error(c *gin.Context, status int, msg string) {
	resp := Response{
		Message: msg,
	}

	c.JSON(status, resp)
}
