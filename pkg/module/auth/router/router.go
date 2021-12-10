package router

import (
	"github.com/gin-gonic/gin"
	"github.com/reynldi/modular-go-boilerplate/pkg/module/auth/delivery/rest"
)

type Router struct {
	handler *rest.Handler
}

type MapRouter struct {
	GetUserByEmail gin.HandlerFunc
}

func NewRouter(handler *rest.Handler) *Router {
	return &Router{
		handler,
	}
}

func (r *Router) Register(gin *gin.RouterGroup) {
	gin.POST("/user", r.handler.AuthUserHandler.GetUserByEmail)
}
