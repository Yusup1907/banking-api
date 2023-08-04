package handler

import (
	"net/http"

	"github.com/Yusup1907/banking-api/src/middleware"
	"github.com/Yusup1907/banking-api/src/model"
	"github.com/Yusup1907/banking-api/src/service"
	"github.com/gin-gonic/gin"
)

type AccountHandler struct {
	router         *gin.Engine
	accountService service.AccountService
}

func (h *AccountHandler) AddAccount(c *gin.Context) {
	var account model.Account

	// Parse the request body into the account struct
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Call the service to add the account
	err := h.accountService.AddAccount(&account)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with success message or the added account data
	c.JSON(http.StatusOK, gin.H{
		"message": "Account added successfully",
	})
}

func NewAccountHandler(r *gin.Engine, accountService service.AccountService) *AccountHandler {
	handler := AccountHandler{
		router:         r,
		accountService: accountService,
	}
	r.POST("/account", middleware.RequireToken(), handler.AddAccount)

	return &handler
}
