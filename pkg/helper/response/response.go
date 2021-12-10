package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response interface {
	OK(message string, data interface{})
}

type response struct {
	c *gin.Context
}

func SendResponse(c *gin.Context) *response {
	return &response{c}
}

func (r *response) OK(message string, data interface{}) {
	formatResponse := FormatSuccessResponse(http.StatusOK, message, data)
	r.c.JSON(http.StatusOK, formatResponse)
}
