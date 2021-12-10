package router

import (
	"github.com/gin-gonic/gin"
	"github.com/reynldi/modular-go-boilerplate/pkg/delivery/rest"
)

type Router struct{}

func SetupRouter(router *gin.Engine, handler *rest.Handler) {
	router.GET("/health", handler.HealthCheck)
}
