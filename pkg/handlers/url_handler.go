package handlers

import (
	"net/http"

	"acheisuacara.com.br/pkg/services"
	"github.com/gin-gonic/gin"
)

type URLHandler struct {
	urlService *services.URLService
}

func NewURLHandler(urlService *services.URLService) *URLHandler {
	return &URLHandler{urlService: urlService}
}

type CreateURLRequest struct {
	URL string `json:"url" binding:"required"`
}

func (h *URLHandler) CreateShortURL(c *gin.Context) {
	var req CreateURLRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	url, err := h.urlService.CreateShortURL(req.URL)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, url)
}

func (h *URLHandler) RedirectToLongURL(c *gin.Context) {
	shortCode := c.Param("shortCode")

	url, err := h.urlService.GetLongURL(shortCode)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	c.Redirect(http.StatusMovedPermanently, url.URL)
}
