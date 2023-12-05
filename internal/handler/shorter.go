package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	urlshorter "github.com/stil4004/url-shorter"
)

func (h *Handler) cutURL(c *gin.Context) {
	var input urlshorter.ShortURL

	if err := c.BindJSON(&input); err != nil{
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	alias_temp, err := h.services.ShorterURL.CreateShortURL(urlshorter.ShortURL{
		Long_url: "abab",
		Short_url: "ab",   
	   })
	
	if err != nil{
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"alias": alias_temp,
	})
}

func (h *Handler) getURL(c *gin.Context) {
	//a := c.Query("url")
	a := c.Param("alias")
	c.JSON(http.StatusOK, map[string]interface{}{
		"allo": a,
	})
}