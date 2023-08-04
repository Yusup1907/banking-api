package handler

import (
	"net/http"
	"strconv"

	"github.com/Yusup1907/banking-api/src/middleware"
	"github.com/Yusup1907/banking-api/src/model"
	"github.com/Yusup1907/banking-api/src/service"
	"github.com/gin-gonic/gin"
)

type NasabahHandler struct {
	router  *gin.Engine
	service service.NasabahService
}

func (h *NasabahHandler) GetAllNasabah(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid page number",
		})
		return
	}

	pageSizeStr := c.DefaultQuery("pageSize", "10")
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid page size",
		})
		return
	}

	nasabahs, err := h.service.GetAllNasabah(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": nasabahs,
	})
}

func (h *NasabahHandler) GetNasabahById(c *gin.Context) {

	id := c.Param("id")

	nasabah, err := h.service.GetNasabahById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if nasabah == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": nasabah,
	})
}

func (h *NasabahHandler) UpdateNasabah(c *gin.Context) {
	var nasabah model.Nasabah

	// Parse the request body into the nasabah struct
	if err := c.ShouldBindJSON(&nasabah); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Call the service to update the nasabah
	err := h.service.UpdateNasabah(&nasabah)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Nasabah not found"})
		return
	}

	// Respond with success message
	c.JSON(http.StatusOK, gin.H{
		"message": "Nasabah updated successfully",
	})
}

func NewNasabahHandler(r *gin.Engine, service service.NasabahService) *NasabahHandler {
	handler := NasabahHandler{
		router:  r,
		service: service,
	}

	r.GET("/nasabah", middleware.RequireToken(), handler.GetAllNasabah)
	r.GET("/nasabah/:id", middleware.RequireToken(), handler.GetNasabahById)
	r.PUT("/nasabah/:id", middleware.RequireToken(), handler.UpdateNasabah)

	return &handler
}
