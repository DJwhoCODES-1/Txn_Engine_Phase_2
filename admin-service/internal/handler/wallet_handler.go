package handler

import (
	"net/http"
	"txn-engine-phase-2/admin-service/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	walletService *service.WalletService
}

func NewHandler(ws *service.WalletService) *Handler {
	return &Handler{
		walletService: ws,
	}
}

func (h *Handler) TopUpWallet(c *gin.Context) {
	var req struct {
		Merchant map[string]interface{} `json:"merchant"`
		Amount   float64                `json:"amount"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// auth skipped (as requested)

	txn, err := h.walletService.TopUpWallet(c, req.Merchant, req.Amount, map[string]interface{}{
		"userId": "admin-id",
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":        true,
		"message":        "Wallet topped up successfully",
		"transaction_id": txn.TransactionID,
		"prevBalance":    txn.PrevBalance,
		"updatedBalance": txn.UpdatedBalance,
	})
}
