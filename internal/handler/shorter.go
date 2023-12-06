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

	alias_temp, err := h.services.ShorterURL.CreateShortURL(&input)
	
	if err != nil{
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"alias": alias_temp,
	})
}

func (h *Handler) getURL(c *gin.Context) {
	var input urlshorter.ShortURL
	long_url := ""

	a := c.Param("alias")
	if a != ""{
		input.Short_url = a
		long_url, err := h.services.ShorterURL.GetLongURL(&input)
		if err != nil || long_url == ""{
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		c.JSON(http.StatusOK, map[string]interface{}{
			"long_url": long_url,
		})
		return
	}

	if err := c.BindJSON(&input); err != nil{
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	long_url, err := h.services.ShorterURL.GetLongURL(&input)

	if err != nil{
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"long_url": long_url,
	})

	}

	