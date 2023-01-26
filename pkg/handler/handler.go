package handler

import (
	"github.com/gin-gonic/gin"
	"qsoft-golang-entering-task/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	timeDifference := router.Group("/when", h.checkHeader)
	{
		timeDifference.GET("/:year", h.getTimeDifference)
	}

	return router
}
