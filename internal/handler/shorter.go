package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	urlshorter "github.com/stil4004/url-shorter"
)

// Function on hadler level, that giving services long url
func (h *Handler) cutURL(c *gin.Context) {

	// Creating var for inputing url
	var input urlshorter.ShortURL

	if err := c.BindJSON(&input); err != nil{
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// Sending to service our long url
	alias_temp, err := h.services.ShorterURL.CreateShortURL(&input)
	
	if err != nil{
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"short_url": alias_temp,
	})
}

func (h *Handler) getURL(c *gin.Context) {

	// Creating var for inputing short url
	var input urlshorter.ShortURL

	// Initing temp var for using it in construction
	long_url := ""

	// Checking if get method was done by adding query params
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

	// Sending short url to services
	long_url, err := h.services.ShorterURL.GetLongURL(&input)

	if err != nil{
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"long_url": long_url,
	})

	}

	