package handler

import (
	"net/http"

	"github.com/Yusup1907/banking-api/src/model"
	"github.com/Yusup1907/banking-api/src/service"
	"github.com/gin-gonic/gin"
)

type MerchantHandler struct {
	router          *gin.Engine
	merchantService service.MerchantService
}

func (h *MerchantHandler) CreateMerchant(c *gin.Context) {
	var merchant model.Merchant

	if err := c.ShouldBindJSON(&merchant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := h.merchantService.AddMerchant(&merchant); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Merchant added successfully",
		"data":    merchant,
	})
}

func (h *MerchantHandler) GetAllMerchants(c *gin.Context) {
	merchants, err := h.merchantService.GetAllMerchants()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": merchants,
	})
}

func (h *MerchantHandler) GetMerchantByID(c *gin.Context) {
	id := c.Param("id")

	merchant, err := h.merchantService.GetMerchantByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if merchant == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Merchant not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": merchant,
	})
}

func (h *MerchantHandler) UpdateMerchant(c *gin.Context) {
	id := c.Param("id")
	var merchant model.Merchant

	if err := c.ShouldBindJSON(&merchant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	merchant.Id = id
	err := h.merchantService.UpdateMerchant(&merchant)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Merchant not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Merchant updated successfully",
		"data":    merchant,
	})
}

func (h *MerchantHandler) DeleteMerchant(c *gin.Context) {
	id := c.Param("id")

	err := h.merchantService.DeleteMerchant(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Merchant not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Merchant deleted successfully",
	})
}

func NewMerchantHandler(r *gin.Engine, merchantService service.MerchantService) *MerchantHandler {
	handler := MerchantHandler{
		router:          r,
		merchantService: merchantService,
	}
	r.POST("/merchant", handler.CreateMerchant)
	r.GET("/merchant", handler.GetAllMerchants)
	r.GET("/merchant/:id", handler.GetMerchantByID)
	r.PUT("/merchant/:id", handler.UpdateMerchant)
	r.DELETE("/merchant", handler.DeleteMerchant)

	return &handler
}
