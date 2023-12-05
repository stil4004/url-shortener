package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/stil4004/url-shorter/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(service *service.Service) *Handler{
	return &Handler{services: service}
}

func (h *Handler) InitRoutes() *gin.Engine{
	router := gin.New()

	api := router.Group("/api")
	{
		api.POST("/shorter", h.cutURL)
		api.GET("/shorter/:alias", h.getURL)
	}
	return router
}