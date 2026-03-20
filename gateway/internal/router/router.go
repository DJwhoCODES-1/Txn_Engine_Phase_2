package router

import (
	"txn-engine-phase-2/gateway/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, authHandler *handler.AuthHandler, walletHandler *handler.WalletHandler) {
	api := r.Group("/api")

	auth := api.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
		auth.POST("/verify-otp", authHandler.VerifyOtp)
	}

	wallet := api.Group("/wallet")
	{
		wallet.POST("/topup", walletHandler.TopUpWallet)
	}

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
}
