package router

import (
	"txn-engine-phase-2/admin-service/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(h *handler.Handler) *gin.Engine {
	r := gin.Default()

	r.POST("/wallet/topup", h.TopUpWallet)

	return r
}
