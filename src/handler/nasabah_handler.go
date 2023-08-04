package handler

import (
	"net/http"
	"strconv"

	"github.com/Yusup1907/banking-api/src/middleware"
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
	id := vars["id"]

	nasabah, err := h.service.GetNasabahById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if nasabah == nil {
		http.Error(w, "Nasabah not found", http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": nasabah,
	})
}

func NewNasabahHandler(r *gin.Engine, service service.NasabahService) *NasabahHandler {
	handler := NasabahHandler{
		router:  r,
		service: service,
	}

	r.GET("/nasabah", middleware.RequireToken(), handler.GetAllNasabah)
	r.GET("/nasabah/{id}", middleware.RequireToken(), handler.GetNasabahById)

	return &handler
}
