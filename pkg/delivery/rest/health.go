package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/reynldi/modular-go-boilerplate/pkg/application/config"
	"github.com/reynldi/modular-go-boilerplate/pkg/helper"
	"github.com/reynldi/modular-go-boilerplate/pkg/helper/response"
)

func (h *Handler) HealthCheck(c *gin.Context) {

	type dataType struct {
		ServiceName string `json:"service_name"`
		Uptime string `json:"uptime"`
	}

	data := dataType{
		ServiceName: config.GetConfig().Application.AppName,
		Uptime: helper.GetUptime(),
	}

	response.SendResponse(c).OK("Success", data)
}
