package handler

import (
	"net/http"

	"txn-engine-phase-2/gateway/internal/client"

	"github.com/gin-gonic/gin"
)

type WalletHandler struct {
	adminClient *client.AdminClient
}

func NewWalletHandler(c *client.AdminClient) *WalletHandler {
	return &WalletHandler{adminClient: c}
}

func (h *WalletHandler) TopUpWallet(c *gin.Context) {
	var req struct {
		Merchant map[string]interface{} `json:"merchant" binding:"required"`
		Amount   float64                `json:"amount" binding:"required"`
		Admin    map[string]interface{} `json:"admin" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	res, err := h.adminClient.TopUpWallet(
		c.Request.Context(),
		req.Merchant,
		req.Amount,
		req.Admin,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
