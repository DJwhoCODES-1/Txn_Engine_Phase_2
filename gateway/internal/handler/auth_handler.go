package handler

import (
	"net/http"

	"txn-engine-phase-2/gateway/internal/client"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	adminClient *client.AdminClient
}

func NewAuthHandler(c *client.AdminClient) *AuthHandler {
	return &AuthHandler{adminClient: c}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req struct {
		Mobile string `json:"mobile" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "mobile required"})
		return
	}

	res, err := h.adminClient.Login(c.Request.Context(), req.Mobile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *AuthHandler) VerifyOtp(c *gin.Context) {
	var req struct {
		Mobile string `json:"mobile" binding:"required"`
		Otp    string `json:"otp" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "mobile & otp required"})
		return
	}

	res, err := h.adminClient.VerifyOtp(c.Request.Context(), req.Mobile, req.Otp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false})
		return
	}

	c.JSON(http.StatusOK, res)
}
