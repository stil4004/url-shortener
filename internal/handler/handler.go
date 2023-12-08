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

	// init main rout
	api := router.Group("/api")
	{
		// rout for posting long url
		api.POST("/shorter", h.cutURL)

		// rout for getting url by short alias
		api.GET("/shorter", h.getURL)
		
		api.GET("/shorter/:alias", h.getURL)
	}
	return router
}