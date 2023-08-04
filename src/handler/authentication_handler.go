package handler

import (
	"log"
	"net/http"
	"time"

	"github.com/Yusup1907/banking-api/src/model"
	"github.com/Yusup1907/banking-api/src/service"
	"github.com/Yusup1907/banking-api/src/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type AuthenticationHandler struct {
	router  *gin.Engine
	service service.NasabahService
}

func (h *AuthenticationHandler) RegisterNasabah(c *gin.Context) {
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

func (h *AuthenticationHandler) Login(c *gin.Context) {
	var loginRequest model.LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	nasabah, err := h.service.Login(&loginRequest, c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Generate a new JWT token
	token, err := utils.GenerateToken(nasabah.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to generate token"})
		return
	}

	// Save the token in the session

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    token,
		HttpOnly: true,
	})

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}

func (h *AuthenticationHandler) Logout(c *gin.Context) {
	// Hapus informasi otentikasi dari sesi
	session := sessions.Default(c)
	session.Clear()
	session.Save()

	// Hapus juga cookie yang menyimpan token JWT
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "token",
		Value:    "", // Nilai cookie dikosongkan
		Path:     "/",
		Expires:  time.Now().Add(-1 * time.Second), // Waktu kadaluwarsa diatur ke masa lalu
		HttpOnly: true,
	})

	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}

func NewAuthenticationHandler(r *gin.Engine, service service.NasabahService) *AuthenticationHandler {
	handler := AuthenticationHandler{
		router:  r,
		service: service,
	}
	r.POST("/register", handler.RegisterNasabah)
	r.POST("/login", handler.Login)
	r.POST("/logout", handler.Logout)

	return &handler
}
