package handler

import (
	"net/http"
	"time"

	"github.com/Yusup1907/banking-api/src/middleware"
	"github.com/Yusup1907/banking-api/src/model"
	"github.com/Yusup1907/banking-api/src/service"
	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	router             *gin.Engine
	transactionService service.TransactionService
}

func (h *TransactionHandler) Transfer(c *gin.Context) {
	var transaction model.Transaction

	// Parse the request body into the transaction struct
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Set the transaction time to the current time
	transaction.TransactionTime = time.Now()

	// Call the service to perform the transfer
	if err := h.transactionService.Transfer(&transaction); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with success message or updated transaction data
	c.JSON(http.StatusOK, gin.H{
		"message":     "Transfer successful",
		"transaction": transaction,
	})
}

func NewTransactionHandler(r *gin.Engine, transactionService service.TransactionService) *TransactionHandler {
	handler := TransactionHandler{
		router:             r,
		transactionService: transactionService,
	}
	r.POST("/payment", middleware.RequireToken(), handler.Transfer)

	return &handler
}
