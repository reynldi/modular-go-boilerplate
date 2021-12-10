package rest

import (
	"net/http"

	"github.com/reynldi/modular-go-boilerplate/pkg/module/auth/model/input"
	"github.com/reynldi/modular-go-boilerplate/pkg/module/auth/service"

	"github.com/gin-gonic/gin"
)

type AuthUserHandler interface {
	GetUserByEmail(c *gin.Context)
}

type authUserHandler struct {
	service service.Service
}

func NewAuthUserHandler(service *service.Service) *authUserHandler {
	return &authUserHandler{
		service: *service,
	}
}

//TODO: Format Response Template
func (h *authUserHandler) GetUserByEmail(c *gin.Context) {
	var input input.GetUserByEmailInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Email is required"})
		return
	}

	c.JSON(http.StatusOK, "user")
}
