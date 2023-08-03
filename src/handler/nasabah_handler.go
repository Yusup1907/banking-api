package handler

import (
	"log"
	"net/http"

	"github.com/Yusup1907/banking-api/src/model"
	"github.com/Yusup1907/banking-api/src/service"
	"github.com/gin-gonic/gin"
)

type NasabahHandler struct {
	router  *gin.Engine
	service service.NasabahService
}

func (h *NasabahHandler) RegisterNasabah(c *gin.Context) {
	var nasabah model.Nasabah
	if err := c.ShouldBindJSON(&nasabah); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	err := h.service.RegisterNasabah(&nasabah)
	if err != nil {
		log.Println("errornya adalah:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Nasabah registered successfully"})
}

func (h *NasabahHandler) Login(c *gin.Context) {
	var loginRequest model.LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	token, err := h.service.Login(loginRequest.Email, loginRequest.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func NewNasabahHandler(r *gin.Engine, service service.NasabahService) *NasabahHandler {
	handler := NasabahHandler{
		router:  r,
		service: service,
	}
	r.POST("/register", handler.RegisterNasabah)
	r.POST("/login", handler.Login)

	return &handler
}
